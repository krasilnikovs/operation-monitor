package model

import (
	"net/url"

	"github.com/google/uuid"
)

const (
	OperationalStatus operationStatus = "operational"
	DegradatedStatus  operationStatus = "degradated"
)

type ServiceId uuid.UUID

func NewServiceId(id string) (ServiceId, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		return ServiceId{}, err
	}

	return ServiceId(uid), nil
}

func (s ServiceId) String() string {
	return uuid.UUID(s).String()
}

type Url url.URL

func (u Url) String() string {
	url := url.URL(u)

	return url.String()
}

type operationStatus string

func (o operationStatus) String() string {
	return string(o)
}

type Service struct {
	id        ServiceId
	name      string
	status    operationStatus
	reference Url
	endpoint  Endpoint
}

func NewService(id ServiceId, name string, reference Url, status operationStatus, endpoint Endpoint) Service {
	return Service{
		id:        id,
		name:      name,
		status:    status,
		reference: reference,
		endpoint:  endpoint,
	}
}

func (s Service) GetId() ServiceId {
	return s.id
}

func (s Service) IsSameId(id ServiceId) bool {
	return s.id == id
}

func (s Service) GetName() string {
	return s.name
}

func (s Service) GetStatus() operationStatus {
	return s.status
}

func (s Service) GetReference() Url {
	return s.reference
}

func (s Service) GetEndpoint() Endpoint {
	return s.endpoint
}

func (s *Service) Degradate() {
	if s.status == DegradatedStatus {
		return
	}

	s.status = DegradatedStatus
}

func (s *Service) Operate() {
	if s.status == OperationalStatus {
		return
	}

	s.status = OperationalStatus
}
