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

type GetServiceById struct {
	repo        repository.ServiceRepository
	transformer transformer.Service
}

func NewGetServiceById(repo repository.ServiceRepository) GetServiceById {
	return GetServiceById{repo: repo}
}

func (h GetServiceById) Execute(id types.ServiceId) (dto.Service, error) {
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
