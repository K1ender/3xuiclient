package xuiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	token      string
	baseURL    string
}

func New(baseURL string, token string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		token:   token,
		baseURL: baseURL,
	}
}

func (c *Client) GetInbounds(ctx context.Context) (Response[[]Inbound], error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/panel/api/inbounds/list", nil)
	if err != nil {
		return Response[[]Inbound]{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Response[[]Inbound]{}, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response[[]Inbound]{}, fmt.Errorf("failed to get inbounds: %s", resp.Status)
	}

	var res Response[[]Inbound]

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Response[[]Inbound]{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return res, nil
}
