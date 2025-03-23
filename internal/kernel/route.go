package kernel

import "github.com/go-chi/chi/v5"

type RouteRegister interface {
	Register(r *chi.Mux)
}

type DefaultRouteRegister struct {
	routeRegisters []RouteRegister
}

func NewDefaultRouteRegister(routeRegisters []RouteRegister) DefaultRouteRegister {
	return DefaultRouteRegister{routeRegisters: routeRegisters}
}

func (drr DefaultRouteRegister) Register(r *chi.Mux) {
	for _, routeRegister := range drr.routeRegisters {
		routeRegister.Register(r)
	}
}
