package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"krasilnikovs.lv/operation-monitor/internal/kernel"
)

func main() {

	kernel.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("APP_NAME")))
	})

	http.ListenAndServe(":3000", r)
}
