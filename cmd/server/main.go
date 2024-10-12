package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sur1k1/go-metrics/internal/server/config"
	"github.com/sur1k1/go-metrics/internal/server/rest"
	"github.com/sur1k1/go-metrics/internal/server/repository/memstorage"
)

func main() {
	serverAddress := config.ParseFlags()

	s := storage.NewStorage()
	router := chi.NewRouter()



	router.Post("/update/{type}/{metric}/{value}", rest.Update(s))
	router.Get("/value/{type}/{metric}", rest.MetricHandler(s))
	router.Get("/", rest.MetricListHandler(s))

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