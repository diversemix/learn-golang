package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/spf13/viper"

	"farmyard/api"
)

func createRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Compress(5),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/farmyard", api.Routes())
	})

	return router
}

func showRoutes(router *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	err := chi.Walk(router, walkFunc)

	if err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
}

func startServer(router *chi.Mux) {
	bind := fmt.Sprintf("%s:%d", viper.GetString("BindAddress"), viper.GetInt("Port"))
	log.Printf("Starting server on http://%s \n", bind)
	log.Fatal(http.ListenAndServe(bind, router))
}

func main() {
	ReadConfiguration()

	router := createRouter()
	showRoutes(router)
	startServer(router)
}
