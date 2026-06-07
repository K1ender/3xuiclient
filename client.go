package xuiclient

import (
	"bytes"
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

func (c *Client) AddClient(ctx context.Context, request CreateClientRequest) error {
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/panel/api/clients/add", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add client: %s", resp.Status)
	}

	return nil
}

func (c *Client) UpdateClient(ctx context.Context, email string, request UpdateClientRequest) error {
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+"/panel/api/clients/update/"+email, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add client: %s", resp.Status)
	}

	return nil
}
