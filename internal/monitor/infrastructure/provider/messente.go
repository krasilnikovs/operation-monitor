package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/model"
	"krasilnikovs.lv/operation-monitor/internal/monitor/domain/provider"
)

type messenteStatusResponse struct {
	Status struct {
		Description string `json:"description"`
	} `json:"status"`
}

func (r messenteStatusResponse) IsOperational() bool {
	return r.Status.Description == "All Systems Operational"
}

type MessenteUptimeProvider struct {
	client http.Client
}

func NewMessenteUptimeProvider(client http.Client) MessenteUptimeProvider {
	return MessenteUptimeProvider{client: client}
}

func (p MessenteUptimeProvider) Supports(m model.Service) bool {
	return m.GetProvider().IsMessente()
}

func (p MessenteUptimeProvider) IsUp(ctx context.Context, m model.Service) (bool, error) {
	if !p.Supports(m) {
		return false, provider.ErrProviderIsNotSupported
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://status.messente.com/api/v2/status.json", nil)

	if err != nil {
		return false, fmt.Errorf("can not create request: %w", err)
	}

	res, err := p.client.Do(req)

	if err != nil {
		return false, fmt.Errorf("fetch resource failed: %w", err)
	}

	defer res.Body.Close()

	var messenteResponse messenteStatusResponse

	err = json.NewDecoder(res.Body).Decode(&messenteResponse)

	if err != nil {
		return false, fmt.Errorf("can not decode json: %w", err)
	}

	return messenteResponse.IsOperational(), nil
}
