package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sur1k1/go-metrics/internal/server/handlers"
	"github.com/sur1k1/go-metrics/internal/server/storage"
)

func main() {
	s := storage.NewStorage()
	router := chi.NewRouter()

	router.Post("/update/{type}/{metric}/{value}", handlers.UpdateHandler(s))
	router.Get("/value/{type}/{metric}", handlers.MetricHandler(s))
	router.Get("/", handlers.MetricListHandler(s))

	if err := run(router); err != nil{
		panic(err)
	}
}

func run(router *chi.Mux) error {
	srv := &http.Server{
		Addr: `:8080`,
		Handler: router,
	}
	return srv.ListenAndServe()
}