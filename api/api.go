package api

import (
	"log"
	"net/http"

	v1 "imnatraj/expense-tracker/api/v1"
	"imnatraj/expense-tracker/middleware"

	"github.com/go-chi/chi/v5"
)

func Api(port, env string) {
	r := chi.NewRouter()
	r.Use(middleware.MetaData())
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", v1.ApiV1(env))
	})
	log.Fatal(http.ListenAndServe(port, r))
}
