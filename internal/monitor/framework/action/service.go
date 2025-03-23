package action

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"krasilnikovs.lv/operation-monitor/configs"
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/handler"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/repository"
)

func GetServiceById(w http.ResponseWriter, r *http.Request) {
	h := handler.NewGetServiceById(
		repository.NewServiceRepository(configs.GetMonitoringServices()),
	)

	serviceId, err := types.NewServiceId(chi.URLParam(r, "serviceId"))

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
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, os.Getenv("APP_NAME"))
}
