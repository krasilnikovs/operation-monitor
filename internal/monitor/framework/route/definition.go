package route

import (
	"github.com/go-chi/chi/v5"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/action"
)

type ServiceRouteRegister struct {
	getServiceByIdAction action.GetServiceByIdAction
}

func NewServiceRouteRegister(getServiceByIdAction action.GetServiceByIdAction) ServiceRouteRegister {
	return ServiceRouteRegister{getServiceByIdAction: getServiceByIdAction}
}

func (sr ServiceRouteRegister) Register(r *chi.Mux) {
	r.Route("/api/v1/services", func(r chi.Router) {
		r.Get("/{serviceId}", sr.getServiceByIdAction.Invoke)
	})

	r.Get("/", action.Index)
}
