package provider

import (
	"context"
	"time"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/provider"
)

type ChainUptimeProvider struct {
	providers []provider.UptimeProvider
}

func NewChainUptimeProvider(providers []provider.UptimeProvider) ChainUptimeProvider {
	return ChainUptimeProvider{providers: providers}
}

func (p ChainUptimeProvider) Supports(m model.Service) bool {
	for _, provider := range p.providers {
		if provider.Supports(m) {
			return true
		}
	}

	return false
}

func (p ChainUptimeProvider) IsUp(ctx context.Context, m model.Service) (bool, error) {
	for _, provider := range p.providers {
		if !provider.Supports(m) {
			continue
		}

		ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*30))

		defer cancel()

		return provider.IsUp(ctx, m)
	}

	return false, provider.ErrProviderIsNotSupported
}
