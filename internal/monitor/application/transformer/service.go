package transformer

import (
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/dto"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
)

type Service struct{}

func NewServiceTransformer() Service {
	return Service{}
}

func (s Service) ToDto(model model.Service) dto.Service {
	return dto.Service{
		Id:        model.GetId().String(),
		Provider:  model.GetProvider().String(),
		Status:    model.GetStatus().String(),
		Reference: model.GetReference().String(),
	}
}
