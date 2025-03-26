package repository

import (
	"context"
	"slices"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

type ServiceRepository struct {
	data []*model.Service
}

func NewServiceRepository(data []*model.Service) *ServiceRepository {
	return &ServiceRepository{data: data}
}

func (s *ServiceRepository) ById(ctx context.Context, id types.ServiceId) (*model.Service, error) {

	defer ctx.Done()

	for _, value := range s.data {
		if value.IsSameId(id) {
			return value, nil
		}
	}

	return nil, nil
}

func (s *ServiceRepository) FetchAll(ctx context.Context) []*model.Service {
	defer ctx.Done()

	return s.data
}

func (s *ServiceRepository) Save(m *model.Service) {
	for index, service := range s.data {
		if !m.IsSameId(service.GetId()) {
			continue
		}

		s.data = slices.Delete(s.data, index, index+1)
		s.data = append(s.data, m)
	}
}
