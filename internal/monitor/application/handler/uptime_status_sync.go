package handler

import (
	"context"
	"fmt"
	"sync"
	"time"

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
	fmt.Println("SYNC LAUNCH")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))

	defer cancel()

	services := uss.repo.FetchAll(ctx)
	var wg sync.WaitGroup

	for _, service := range services {
		if !uss.uptimeProvider.Supports(service) {
			continue
		}

		wg.Add(1)

		go func() {

			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*30))

			defer cancel()

			isUp, err := uss.uptimeProvider.IsUp(ctx, service)

			if err != nil {
				service.Pending()
				return
			}

			if isUp {
				service.Operate()
			}

			if !isUp {
				service.Degradate()
			}
		}()

		wg.Wait()

		for _, service := range services {
			uss.repo.Save(service)
		}
	}
}
