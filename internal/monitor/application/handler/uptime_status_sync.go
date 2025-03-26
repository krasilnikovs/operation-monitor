package handler

import (
	"context"
	"sync"
	"time"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/provider"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/repository"
)

type UptimeStatusSync struct {
	repo           repository.ServiceRepository
	uptimeProvider provider.UptimeProvider
}

func NewUptimeStatusSync(repo repository.ServiceRepository, uptimeProvider provider.UptimeProvider) UptimeStatusSync {
	return UptimeStatusSync{repo: repo, uptimeProvider: uptimeProvider}
}

func (uss UptimeStatusSync) Execute() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))

	defer cancel()

	services := uss.repo.FetchAll(ctx)
	var wg sync.WaitGroup

	for _, service := range services {
		if !uss.uptimeProvider.Supports(*service) {
			continue
		}

		wg.Add(1)

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))

			defer wg.Done()
			defer cancel()

			uss.sync(ctx, service)
		}()
	}

	wg.Wait()
}

func (uss *UptimeStatusSync) sync(ctx context.Context, service *model.Service) {

	isUp, err := uss.uptimeProvider.IsUp(ctx, *service)

	if err != nil {
		service.Degradate()
		uss.repo.Save(service)
		return
	}

	if isUp {
		service.Operate()
	}

	if !isUp {
		service.Degradate()
	}

	uss.repo.Save(service)
}
