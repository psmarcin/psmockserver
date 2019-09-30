package server

import (
	"net/http"

	"github.com/kataras/golog"
)

func Start(port string) {
	router := CreateRouter()
	golog.Fatal(http.ListenAndServe(":"+port, router))
}
