package route

import (
	"github.com/go-chi/chi/v5"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/action"
)

type ServiceRouteRegister struct{}

func NewServiceRouteRegister() ServiceRouteRegister {
	return ServiceRouteRegister{}
}

func (sr ServiceRouteRegister) Register(r *chi.Mux) {
	r.Route("/api/v1/services", func(r chi.Router) {
		r.Get("/{serviceId}", action.GetServiceById)
	})

	r.Get("/", action.Index)
}
