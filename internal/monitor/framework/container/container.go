package container

import (
	"net/http"
	"time"

	"krasilnikovs.lv/operation-monitor/configs"
	"krasilnikovs.lv/operation-monitor/internal/monitor/application/handler"
	providerContract "krasilnikovs.lv/operation-monitor/internal/monitor/domain/provider"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/action"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/jobs"
	"krasilnikovs.lv/operation-monitor/internal/monitor/framework/route"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/provider"
	"krasilnikovs.lv/operation-monitor/internal/monitor/infrastructure/repository"
	"krasilnikovs.lv/operation-monitor/pkg/kernel"
)

type ServiceContainer struct{}

func NewServiceContainer() ServiceContainer {
	return ServiceContainer{}
}

var jobLoader kernel.JobLoader = nil
var routeRegister kernel.RouteRegister = nil
var getServiceByIdHandler *handler.GetServiceById = nil
var uptimeStatusSyncHandler *handler.UptimeStatusSync = nil
var serviceRepository *repository.ServiceRepository = nil
var getServiceByIdAction *action.GetServiceByIdAction = nil
var serviceRouteRegister *route.ServiceRouteRegister = nil
var chainUptimeProvider *provider.ChainUptimeProvider = nil
var messenteUptimeProvider *provider.MessenteUptimeProvider = nil

func (sc ServiceContainer) ProvideDefaultRouteRegister() kernel.RouteRegister {
	if routeRegister == nil {
		rr := kernel.NewDefaultRouteRegister(sc.provideRouteRegisters())

		routeRegister = rr
	}

	return routeRegister
}

func (sc ServiceContainer) ProvideDefaultJobLoader() kernel.JobLoader {
	if jobLoader == nil {
		jl := kernel.NewDefaultJobLoader(sc.provideJobsLoaders())
		jobLoader = &jl
	}

	return jobLoader
}

func (sc ServiceContainer) provideRouteRegisters() []kernel.RouteRegister {
	return []kernel.RouteRegister{
		sc.provideServiceRouteRegister(),
	}
}

func (sc ServiceContainer) provideNewMessenteUptimeProvider() provider.MessenteUptimeProvider {
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
			sc.provideServiceRepository(),
			sc.provideNewChainUptimeProvider(),
		)

		uptimeStatusSyncHandler = &h
	}

	return *uptimeStatusSyncHandler
}

func (sc ServiceContainer) provideGetServiceByIdHandler() handler.GetServiceById {
	if getServiceByIdHandler == nil {
		h := handler.NewGetServiceById(sc.provideServiceRepository())

		getServiceByIdHandler = &h
	}

	return *getServiceByIdHandler
}

func (sc ServiceContainer) provideServiceRepository() *repository.ServiceRepository {
	if serviceRepository == nil {
		r := repository.NewServiceRepository(configs.GetMonitoringServices())

		serviceRepository = r
	}

	return serviceRepository
}

func (sc ServiceContainer) provideGetServiceByIdAction() action.GetServiceByIdAction {
	if getServiceByIdAction == nil {
		a := action.NewGetServiceByIdAction(sc.provideGetServiceByIdHandler())

		getServiceByIdAction = &a
	}

	return *getServiceByIdAction
}

func (sc ServiceContainer) provideServiceRouteRegister() route.ServiceRouteRegister {
	if serviceRouteRegister == nil {
		s := route.NewServiceRouteRegister(sc.provideGetServiceByIdAction())

		serviceRouteRegister = &s
	}

	return *serviceRouteRegister
}

func (sc ServiceContainer) provideNewChainUptimeProvider() provider.ChainUptimeProvider {
	if chainUptimeProvider == nil {
		p := provider.NewChainUptimeProvider(
			[]providerContract.UptimeProvider{
				sc.provideNewMessenteUptimeProvider(),
			},
		)

		chainUptimeProvider = &p
	}

	return *chainUptimeProvider
}

func (sc ServiceContainer) provideJobsLoaders() []kernel.JobLoader {
	return []kernel.JobLoader{
		jobs.NewUptimeStatusSyncJobLoader(
			sc.ProvideNewUptimeStatusSyncHandler(),
		),
	}
}
