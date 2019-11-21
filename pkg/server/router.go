package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"psmockserver/pkg/mock"
	"psmockserver/pkg/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kataras/golog"
)

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	// middlewares
	router.Use(LogMiddleware)
	router.Use(middleware.Recoverer)

	// healthcheck
	router.Get("/health/r", readinesHandler)
	router.Get("/health/l", livenessHandler)

	// routes
	router.Post(`/mockserver`, addMockHandler)
	router.Get(`/mockserver`, listMockHandler)
	router.Put(`/mockserver/reset`, resetHandler)
	router.HandleFunc(`/*`, rootHandler)
	router.NotFound(http.NotFound)

	return router
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	queryStrings := r.URL.Query()
	reqID := mock.GetMockHash(mock.RequestId{
		Method:       r.Method,
		Path:         r.URL.Path,
		QueryStrings: queryStrings,
	})
	m, err := mock.Find(reqID)
	if err != nil {
		golog.Warnf("Didn't find mock for: %s %s", r.Method, r.RequestURI)
		mock.List()
		http.NotFound(w, r)
		return
	}
	// set headers
	for k, v := range m.Headers {
		w.Header().Set(k, v[0])
	}
	// set statusCode
	w.WriteHeader(m.StatusCode)
	fmt.Fprint(w, m.Body)
}

func listMockHandler(w http.ResponseWriter, r *http.Request) {
	str, err := mock.Serialize()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(str)
}

func addMockHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// todo: handle error on parse
	body, _ := ioutil.ReadAll(r.Body)
	p, err := mock.Parse(body)
	if err != nil {
		http.Error(w, "cant parse", http.StatusBadRequest)
		return
	}
	mock.Add(mock.GetMockHash(mock.RequestId{
		Method:       p.HttpRequest.Method,
		Path:         p.HttpRequest.Path,
		QueryStrings: p.HttpRequest.QueryStrings,
	}), mock.Mock{
		Headers:     utils.AddHeaders(p.HttpResponse.Headers),
		StatusCode:  p.HttpResponse.StatusCode,
		Body:        p.HttpResponse.Body,
		ContentType: p.HttpRequest.ContentType,
		Method:      p.HttpRequest.Method,
		RemainingTimes: mock.Remaining{
			Times:     p.Times.RemainingTimes,
			Unlimited: p.Times.Unlimited,
		},
	})
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	mock.Reset()
	w.WriteHeader(http.StatusAccepted)
}
