package provider

import (
	"context"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
)

type UptimeProvider interface {
	Supports(model.Service) bool
	IsUp(context.Context, model.Service) (bool, error)
}
