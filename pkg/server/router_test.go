package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"psmockserver/pkg/mock"
)

func Test_rootHandler(t *testing.T) {
	t.Run("should return default content-type", func(t *testing.T) {
		response := httptest.NewRecorder()
		path := "/create"
		id := mock.GetMockHash(mock.RequestId{
			Method:       http.MethodGet,
			Path:         path,
			QueryStrings: url.Values{},
		})
		mock.M.Add(id, mock.MockResponse{
			Body: "body",
			RemainingTimes: mock.Remaining{
				Unlimited: true,
			},
			ContentType: "application/json",
		})

		request, _ := http.NewRequest(http.MethodGet, path, nil)
		CreateRouter().ServeHTTP(response, request)
		deb, _ := mock.M.Serialize()
		fmt.Printf("deb %s\n", deb)
		fmt.Printf("contenttype %s", response.Header().Get("content-type"))
		if response.Header().Get("content-type") != "application/json" {
			t.Fatalf("expected to return %s but got %s", "application/json", response.Header().Get("content-type"))
		}
	})
}
