package repository

import (
	"context"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

type ServiceRepository interface {
	ById(context.Context, types.ServiceId) (*model.Service, error)
	FetchAll(context.Context) []model.Service
	Save(model.Service)
}
