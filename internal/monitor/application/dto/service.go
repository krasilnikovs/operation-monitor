package dto

type Service struct {
	Id        string `json:"id"`
	Provider  string `json:"provider"`
	Status    string `json:"status"`
	Reference string `json:"reference"`
}
