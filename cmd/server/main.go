package main

import (
	"net/http"

	"github.com/sur1k1/go-metrics/internal/server/handlers"
	"github.com/sur1k1/go-metrics/internal/server/storage"
)

func main() {
	s := storage.NewStorage()

	mux := http.NewServeMux()
	mux.HandleFunc("/update/{type}/{metric}/{value}", handlers.UpdateHandler(s))

	if err := run(mux); err != nil{
		panic(err)
	}
}

func run(mux *http.ServeMux) error {
	return http.ListenAndServe(`:8080`, mux)
}