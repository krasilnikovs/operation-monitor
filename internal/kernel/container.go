package kernel

import (
	"net/http"
	"time"

	"krasilnikovs.lv/operation-monitor/configs"
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/handler"
	providerContract "krasilnikovs.lv/operation-monitor/internal/monitor/domain/provider"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/action"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/route"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/provider"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/repository"
)

type ServiceContainer struct{}

func NewServiceContainer() ServiceContainer {
	return ServiceContainer{}
}

var getServiceByIdHandler *handler.GetServiceById = nil
var uptimeStatusSyncHandler *handler.UptimeStatusSync = nil
var serviceRepository *repository.ServiceRepository = nil
var getServiceByIdAction *action.GetServiceByIdAction = nil
var serviceRouteRegister *route.ServiceRouteRegister = nil
var chainUptimeProvider *provider.ChainUptimeProvider = nil
var messenteUptimeProvider *provider.MessenteUptimeProvider = nil

func (sc ServiceContainer) ProvideGetServiceByIdHandler() handler.GetServiceById {
	if getServiceByIdHandler == nil {
		h := handler.NewGetServiceById(sc.ProvideServiceRepository())

		getServiceByIdHandler = &h
	}

	return *getServiceByIdHandler
}

func (sc ServiceContainer) ProvideServiceRepository() repository.ServiceRepository {
	if serviceRepository == nil {
		r := repository.NewServiceRepository(configs.GetMonitoringServices())

		serviceRepository = &r
	}

	return *serviceRepository
}

func (sc ServiceContainer) ProvideGetServiceByIdAction() action.GetServiceByIdAction {
	if getServiceByIdAction == nil {
		a := action.NewGetServiceByIdAction(sc.ProvideGetServiceByIdHandler())

		getServiceByIdAction = &a
	}

	return *getServiceByIdAction
}

func (sc ServiceContainer) ProvideServiceRouteRegister() route.ServiceRouteRegister {
	if serviceRouteRegister == nil {
		s := route.NewServiceRouteRegister(sc.ProvideGetServiceByIdAction())

		serviceRouteRegister = &s
	}

	return *serviceRouteRegister
}

func (sc ServiceContainer) ProvideRouteRegisters() []RouteRegister {
	return []RouteRegister{
		sc.ProvideServiceRouteRegister(),
	}
}

func (sc ServiceContainer) ProvideNewChainUptimeProvider() provider.ChainUptimeProvider {
	if chainUptimeProvider == nil {
		p := provider.NewChainUptimeProvider(
			[]providerContract.UptimeProvider{
				sc.ProvideNewMessenteUptimeProvider(),
			},
		)

		chainUptimeProvider = &p
	}

	return *chainUptimeProvider
}

func (sc ServiceContainer) ProvideNewMessenteUptimeProvider() provider.MessenteUptimeProvider {
	if messenteUptimeProvider == nil {
		p := provider.NewMessenteUptimeProvider(http.Client{
			Timeout: time.Duration(30 * time.Second),
		})

		messenteUptimeProvider = &p
	}

	return *messenteUptimeProvider
}

func (sc ServiceContainer) ProvideNewUptimeStatusSyncHandler() handler.UptimeStatusSync {
	if uptimeStatusSyncHandler == nil {
		h := handler.NewUptimeStatusSync(
			sc.ProvideServiceRepository(),
			sc.ProvideNewChainUptimeProvider(),
		)

		uptimeStatusSyncHandler = &h
	}

	return *uptimeStatusSyncHandler
}
