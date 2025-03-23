package kernel

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func Load() {
	loadDotEnv()
}

func LoadWeb(routes []RouteRegister) *chi.Mux {
	loadDotEnv()

	return loadRoutes(routes)
}

func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error during loading .env: %w", err)
	}

	godotenv.Overload(".env.local")
}

func loadRoutes(routes []RouteRegister) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	routeRegister := NewDefaultRouteRegister(routes)

	routeRegister.Register(r)

	return r
}
