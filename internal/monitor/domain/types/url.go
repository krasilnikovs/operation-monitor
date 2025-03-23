package types

import (
	"fmt"
	"net/url"
)

type Url url.URL

func NewUrl(value string) (Url, error) {
	parsedUrl, err := url.Parse(value)

	if err != nil {
		return Url{}, fmt.Errorf("invalid url")
	}

	return Url(*parsedUrl), nil
}

func (u Url) String() string {
	url := url.URL(u)

	return url.String()
}
