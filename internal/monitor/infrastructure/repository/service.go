package repository

import (
	"context"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

type Service struct {
	data []model.Service
}

func NewServiceRepository(data []model.Service) Service {
	return Service{data: data}
}

func (s Service) ById(ctx context.Context, id types.ServiceId) (*model.Service, error) {

	defer ctx.Done()

	for _, value := range s.data {
		if value.IsSameId(id) {
			return &value, nil
		}
	}

	return nil, nil
}
