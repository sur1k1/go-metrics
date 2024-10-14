package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sur1k1/go-metrics/internal/server/config"
	"github.com/sur1k1/go-metrics/internal/server/repository/memstorage"
	"github.com/sur1k1/go-metrics/internal/server/rest"
	"github.com/sur1k1/go-metrics/service"
)

func main() {
	serverAddress := config.ParseFlags()

	// Prepare chi
	router := chi.NewRouter()

	// Prepare repository
	s := storage.NewStorage()

	// Build service layer
	svc := service.NewService(s)
	rest.NewUpdateHandler(router, svc)
	rest.NewMetricHandler(router, svc)
	rest.NewMetricListHandler(router, svc)

	// Start server
	if err := run(router, serverAddress); err != nil{
		fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
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