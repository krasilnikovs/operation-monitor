package configs

import (
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

func GetMonitoringServices() []*model.Service {
	return []*model.Service{
		messente(),
	}
}

func messente() *model.Service {
	id, _ := types.NewServiceId("0195bf98-9f36-7a71-91e3-ded76ada3edb")
	reference, _ := types.NewUrl("https://messente.com")

	s := model.NewService(
		id,
		types.MessenteServiceProvider,
		reference,
		types.PendingStatus,
	)

	return &s
}
