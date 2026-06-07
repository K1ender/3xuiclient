package xuiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (c *Client) do(
	ctx context.Context,
	method string,
	path string,
	body any,
	out any,
) error {
	var reader *bytes.Reader

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal request: %w", err)
		}

		reader = bytes.NewReader(data)
	} else {
		reader = bytes.NewReader(nil)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		c.baseURL+path,
		reader,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) GetInbounds(ctx context.Context) (Response[[]Inbound], error) {
	var res Response[[]Inbound]

	err := c.do(
		ctx,
		http.MethodGet,
		"/panel/api/inbounds/list",
		nil,
		&res,
	)

	return res, err
}

func (c *Client) AddClient(ctx context.Context, req CreateClientRequest) error {
	return c.do(
		ctx,
		http.MethodPost,
		"/panel/api/clients/add",
		req,
		nil,
	)
}

func (c *Client) UpdateClient(
	ctx context.Context,
	email string,
	req UpdateClientRequest,
) error {
	return c.do(
		ctx,
		http.MethodPost,
		"/panel/api/clients/update/"+url.PathEscape(email),
		req,
		nil,
	)
}
