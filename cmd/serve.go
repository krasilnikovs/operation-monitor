package main

import (
	"net/http"

	"krasilnikovs.lv/operation-monitor/internal/kernel"
	monitorRoute "krasilnikovs.lv/operation-monitor/internal/monitor/framework/route"
)

var routeRegisters = []kernel.RouteRegister{
	monitorRoute.NewServiceRouteRegister(),
}

func main() {
	r := kernel.LoadWeb(routeRegisters)

	http.ListenAndServe(":3000", r)
}
