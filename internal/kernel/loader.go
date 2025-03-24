package kernel

import (
	"log"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

var serviceContainer *ServiceContainer = nil

func Load() {
	loadDotEnv()
}

func LoadWeb() *chi.Mux {
	loadDotEnv()
	loadServiceContainer()
	loadJobs()

	return loadRoutes()
}

func loadDotEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("error during loading .env: %w", err)
	}

	godotenv.Overload(".env.local")
}

func loadServiceContainer() {
	if serviceContainer == nil {
		conrainer := NewServiceContainer()
		serviceContainer = &conrainer
	}
}

func loadRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	routeRegister := NewDefaultRouteRegister(serviceContainer.ProvideRouteRegisters())

	routeRegister.Register(r)

	return r
}

func loadJobs() {
	h := serviceContainer.ProvideNewUptimeStatusSyncHandler()

	ticker := time.NewTicker(time.Duration(time.Second * 3))

	go func() {
		for {
			<-ticker.C

			h.Execute()
		}
	}()
}
