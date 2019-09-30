package server

import (
	"net/http"

	"github.com/kataras/golog"
)

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		golog.Infof("%s %s", r.Method, r.URL.String())
		h.ServeHTTP(w, r)
	})
}
