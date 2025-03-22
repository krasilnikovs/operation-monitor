package repository

import (
	"context"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
)

type Service interface {
	ById(ctx context.Context, id model.ServiceId) (*model.Service, error)
}
