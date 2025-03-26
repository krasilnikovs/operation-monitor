package jobs

import (
	"time"

	"krasilnikovs.lv/operation-monitor/internal/monitor/application/handler"
)

type UptimeStatusSyncJobLoader struct {
	h handler.UptimeStatusSync
}

func NewUptimeStatusSyncJobLoader(h handler.UptimeStatusSync) UptimeStatusSyncJobLoader {
	return UptimeStatusSyncJobLoader{h: h}
}

func (j UptimeStatusSyncJobLoader) LoadJob() {
	ticker := time.NewTicker(time.Duration(time.Second * 30))

	go func() {
		for {
			<-ticker.C

			j.h.Execute()
		}
	}()
}
