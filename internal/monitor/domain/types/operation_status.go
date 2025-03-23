package types

const (
	OperationalStatus OperationStatus = "operational"
	PendingStatus     OperationStatus = "pending"
	DegradatedStatus  OperationStatus = "degradated"
)

type OperationStatus string

func (o OperationStatus) String() string {
	return string(o)
}
