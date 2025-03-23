package handler

import (
	"context"
	"fmt"
	"time"

	"krasilnikovs.lv/operation-monitor/internal/monitor/application/dto"
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/transformer"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/repository"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

var (
	ErrServiceNotFound = fmt.Errorf("service is not found")
)

type getServiceById struct {
	repo        repository.Service
	transformer transformer.Service
}

func NewGetServiceById(repo repository.Service) getServiceById {
	return getServiceById{repo: repo}
}

func (h getServiceById) Execute(id types.ServiceId) (dto.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))

	defer cancel()

	service, err := h.repo.ById(ctx, id)

	if service == nil {
		return dto.Service{}, ErrServiceNotFound
	}

	if err != nil {
		return dto.Service{}, err
	}

	return h.transformer.ToDto(*service), nil
}
