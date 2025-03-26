package kernel

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type Kernel struct {
	routeRegister RouteRegister
	jobLoader     JobLoader
}

func NewKernel(routeRegister RouteRegister, jobLoader JobLoader) Kernel {
	return Kernel{routeRegister: routeRegister, jobLoader: jobLoader}
}

func (k Kernel) LoadWeb() http.Handler {
	k.loadDotEnv()

	defer k.jobLoader.LoadJob()

	return k.loadRoutes()
}

func (k Kernel) loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error during loading .env: %w", err)
	}

	godotenv.Overload(".env.local")
}

func (k Kernel) loadRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	k.routeRegister.Register(r)

	return r
}
