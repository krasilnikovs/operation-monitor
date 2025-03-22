package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"krasilnikovs.lv/operation-monitor/internal/kernel"
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/handler"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/repository"
)

func main() {

	kernel.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(os.Getenv("APP_NAME")))
	})

	r.Route("/api/v1/services", func(r chi.Router) {
		r.Get("/{serviceId}", func(w http.ResponseWriter, r *http.Request) {
			h := handler.NewGetServiceById(
				repository.NewServiceRepository([]model.Service{}),
			)

			serviceId, err := model.NewServiceId(chi.URLParam(r, "serviceId"))

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(err)
				return
			}

			dto, err := h.Execute(serviceId)

			if err == handler.ErrServiceNotFound {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err)
				return
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(err)
				return
			}

			json.NewEncoder(w).Encode(dto)
		})
	})

	http.ListenAndServe(":3000", r)
}
