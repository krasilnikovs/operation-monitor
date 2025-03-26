package main

import (
	"net/http"

	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/container"
	"krasilnikovs.lv/operation-monitor/pkg/kernel"
)

func main() {
	serviceContainer := container.NewServiceContainer()

	k := kernel.NewKernel(
		serviceContainer.ProvideDefaultRouteRegister(),
		serviceContainer.ProvideDefaultJobLoader(),
	)

	r := k.LoadWeb()

	http.ListenAndServe(":3000", r)
}
