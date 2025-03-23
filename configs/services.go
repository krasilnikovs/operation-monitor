package configs

import (
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
)

func GetMonitoringServices() []model.Service {
	return []model.Service{
		messente(),
	}
}

func messente() model.Service {
	id, _ := model.NewServiceId("0195bf98-9f36-7a71-91e3-ded76ada3edb")
	reference, _ := model.NewUrl("https://messente.com")
	uri, _ := model.NewUrl("https://status.messente.com/api/v2/status.json")

	return model.NewService(
		id,
		"Messente",
		reference,
		model.PendingStatus,
		model.NewEndpoint(
			uri,
			model.ApplicationJsonType,
		),
	)
}
