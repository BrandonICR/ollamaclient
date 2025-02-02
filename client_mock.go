package ollamaclient

import (
	"context"
	"net/http"
)

// ClientMock mock implementation of Client interface.
type ClientMock struct {
	GenerateFunc func(ctx context.Context, req GenerateRequest) (GenerateResponse, error)
	ChatFunc     func(ctx context.Context, req ChatRequest) (GenerateResponse, error)
}

// Generate mock implementation of Client.Generate interface method.
func (m *ClientMock) Generate(ctx context.Context, req GenerateRequest) (GenerateResponse, error) {
	return m.GenerateFunc(ctx, req)
}

// Chat mock implementation of Client.Chat interface method.
func (m *ClientMock) Chat(ctx context.Context, req ChatRequest) (GenerateResponse, error) {
	return m.ChatFunc(ctx, req)
}

// HTTPClientMock mock implementation of HTTPClient interface.
type HTTPClientMock struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do mock implementation of HTTPClient.Do interface method.
func (m *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
