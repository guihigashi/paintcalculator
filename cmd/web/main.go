package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/v1/cans-needed", SolveCansNeeded)

	port := 8080
	fmt.Printf("serving at port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
