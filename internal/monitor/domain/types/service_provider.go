package types

const (
	MessenteServiceProvider ServiceProvider = "messente"
)

type ServiceProvider string

func (sp ServiceProvider) String() string {
	return string(sp)
}

func (sp ServiceProvider) IsMessente() bool {
	return sp == MessenteServiceProvider
}
