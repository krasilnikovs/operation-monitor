package model

import (
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/types"
)

type Service struct {
	id        types.ServiceId
	provider  types.ServiceProvider
	status    types.OperationStatus
	reference types.Url
}

func NewService(id types.ServiceId, provider types.ServiceProvider, reference types.Url, status types.OperationStatus) Service {
	return Service{
		id:        id,
		provider:  provider,
		status:    status,
		reference: reference,
	}
}

func (s Service) GetId() types.ServiceId {
	return s.id
}

func (s Service) IsSameId(id types.ServiceId) bool {
	return s.id == id
}

func (s Service) GetProvider() types.ServiceProvider {
	return s.provider
}

func (s Service) GetStatus() types.OperationStatus {
	return s.status
}

func (s Service) GetReference() types.Url {
	return s.reference
}

func (s *Service) Degradate() {
	if s.status == types.DegradatedStatus {
		return
	}

	s.status = types.DegradatedStatus
}

func (s *Service) Operate() {
	if s.status == types.OperationalStatus {
		return
	}

	s.status = types.OperationalStatus
}
