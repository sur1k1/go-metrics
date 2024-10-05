package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sur1k1/go-metrics/internal/server/config"
	"github.com/sur1k1/go-metrics/internal/server/handlers"
	"github.com/sur1k1/go-metrics/internal/server/storage"
)

func main() {
	serverAddress := config.ParseFlags()

	s := storage.NewStorage()
	router := chi.NewRouter()

	router.Post("/update/{type}/{metric}/{value}", handlers.UpdateHandler(s))
	router.Get("/value/{type}/{metric}", handlers.MetricHandler(s))
	router.Get("/", handlers.MetricListHandler(s))

	if err := run(router, serverAddress); err != nil{
		panic(err)
	}
}

func run(router *chi.Mux, serverAddress string) error {
	srv := &http.Server{
		Addr: serverAddress,
		Handler: router,
	}
	fmt.Println("Server started:", serverAddress)
	return srv.ListenAndServe()
}