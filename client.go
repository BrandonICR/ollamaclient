package ollamaclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	Generate(ctx context.Context, req GenerateRequest) (GenerateResponse, error)
	Chat(ctx context.Context, req ChatRequest) (ChatResponse, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// client implementation of Client interface.
type client struct {
	domain string
	client HTTPClient
}

func NewClient(c HTTPClient, domain string) Client {
	return &client{client: c, domain: domain}
}

// Generate implementation of Client.Generate interface method.
func (c *client) Generate(ctx context.Context, req GenerateRequest) (GenerateResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return GenerateResponse{}, fmt.Errorf("marshal request: %v", err)
	}
	resBody, err := c.do(ctx, "/api/generate", reqBody)
	if err != nil {
		return GenerateResponse{}, err
	}
	var res GenerateResponse
	if err := json.Unmarshal(resBody, &res); err != nil {
		return GenerateResponse{}, fmt.Errorf("unmarshal response: %v", err)
	}
	return res, nil
}

// Chat implementation of Client.Chat interface method.
func (c *client) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("marshal request: %v", err)
	}
	resBody, err := c.do(ctx, "/api/chat", reqBody)
	if err != nil {
		return ChatResponse{}, err
	}
	var res ChatResponse
	if err := json.Unmarshal(resBody, &res); err != nil {
		return ChatResponse{}, fmt.Errorf("unmarshal response: %v", err)
	}
	return res, nil
}

func (c *client) do(ctx context.Context, path string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.domain+path, io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status_code [%d] with body [%s]", resp.StatusCode, string(respBody))
	}
	return respBody, nil
}
