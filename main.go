package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		_, _ = fmt.Fprint(writer, "Get")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	_ = server.ListenAndServe()
}
