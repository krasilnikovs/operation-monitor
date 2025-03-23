package repository

import (
	"context"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

type Service interface {
	ById(ctx context.Context, id types.ServiceId) (*model.Service, error)
}
