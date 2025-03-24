package provider

import (
	"context"
	"fmt"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
)

var (
	ErrProviderIsNotSupported = fmt.Errorf("provider not supported")
)

type UptimeProvider interface {
	Supports(model.Service) bool
	IsUp(context.Context, model.Service) (bool, error)
}
