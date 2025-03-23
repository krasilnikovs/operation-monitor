package types

import "github.com/google/uuid"

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
