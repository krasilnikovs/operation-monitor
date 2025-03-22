package model

const (
	ApplicationJsonType contentType = "application/json"
)

type contentType string

type Endpoint struct {
	uri         Url
	contentType contentType
}

func NewEndpoint(uri Url, contentType contentType) Endpoint {
	return Endpoint{uri: uri, contentType: contentType}
}
