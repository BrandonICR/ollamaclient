package ollamaclient_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/BrandonICR/ollamaclient"
)

func TestClient_Generate(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
					"model": "llama3.2",
					"created_at": "2023-08-04T19:22:45.499127Z",
					"response": "The sky is blue because it is the color of the sky.",
					"done": true,
					"context": [1, 2, 3],
					"total_duration": 5043500667,
					"load_duration": 5025959,
					"prompt_eval_count": 26,
					"prompt_eval_duration": 325953000,
					"eval_count": 290,
					"eval_duration": 4709213000
				}`)),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewGenerateRequest("llama3.2", "Why is the sky blue?").WithStream(false).Build()
	expected := ollamaclient.GenerateResponse{
		Model:              "llama3.2",
		CreatedAt:          "2023-08-04T19:22:45.499127Z",
		Response:           "The sky is blue because it is the color of the sky.",
		Done:               true,
		Context:            []int{1, 2, 3},
		TotalDuration:      5043500667,
		LoadDuration:       5025959,
		PromptEvalCount:    26,
		PromptEvalDuration: 325953000,
		EvalCount:          290,
		EvalDuration:       4709213000,
	}

	res, err := client.Generate(ctx, req)

	if err != nil {
		t.Fatalf("something was wrong: %v", err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
}

func TestClient_Generate_ErrorUnmarshalResponseBody(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`["unexpected_body"]`)),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewGenerateRequest("llama3.2", "Why is the sky blue?").WithStream(false).Build()
	expected := ollamaclient.GenerateResponse{}

	res, err := client.Generate(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "unmarshal response:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Generate_ErrorUnexpectedStatusCode(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewGenerateRequest("llama3.2", "Why is the sky blue?").WithStream(false).Build()
	expected := ollamaclient.GenerateResponse{}

	res, err := client.Generate(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "status_code [404] with body") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Generate_ErrorDoRequest(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("something went wrong")
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewGenerateRequest("llama3.2", "Why is the sky blue?").WithStream(false).Build()
	expected := ollamaclient.GenerateResponse{}

	res, err := client.Generate(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "do request:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Generate_ErrorCreatingRequest(t *testing.T) {
	ctx := context.Context(nil)
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("something went wrong")
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewGenerateRequest("llama3.2", "Why is the sky blue?").WithStream(false).Build()
	expected := ollamaclient.GenerateResponse{}

	res, err := client.Generate(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "creating request:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Chat(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body: io.NopCloser(strings.NewReader(`{
					"model": "llama3.2",
					"created_at": "2023-12-12T14:13:43.416799Z",
					"message": {
						"role": "assistant",
						"content": "Hello! How are you today?"
					},
					"done": true,
					"total_duration": 5191566416,
					"load_duration": 2154458,
					"prompt_eval_count": 26,
					"prompt_eval_duration": 383809000,
					"eval_count": 298,
					"eval_duration": 4799921000
				}`)),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewChatRequest("llama3.2", []ollamaclient.Message{{Role: "user", Content: "why is the sky blue?"}}).
		WithStream(false).Build()
	expected := ollamaclient.ChatResponse{
		Model:     "llama3.2",
		CreatedAt: "2023-12-12T14:13:43.416799Z",
		Message: ollamaclient.Message{
			Role:    "assistant",
			Content: "Hello! How are you today?",
		},
		Done:               true,
		TotalDuration:      5191566416,
		LoadDuration:       2154458,
		PromptEvalCount:    26,
		PromptEvalDuration: 383809000,
		EvalCount:          298,
		EvalDuration:       4799921000,
	}

	res, err := client.Chat(ctx, req)

	if err != nil {
		t.Fatalf("something was wrong: %v", err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
}

func TestClient_Chat_ErrorUnmarshalResponseBody(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`["unexpected_body"]`)),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewChatRequest("llama3.2", []ollamaclient.Message{{Role: "user", Content: "why is the sky blue?"}}).
		WithStream(false).Build()
	expected := ollamaclient.ChatResponse{}

	res, err := client.Chat(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "unmarshal response:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Chat_ErrorUnexpectedStatusCode(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewChatRequest("llama3.2", []ollamaclient.Message{{Role: "user", Content: "why is the sky blue?"}}).
		WithStream(false).Build()
	expected := ollamaclient.ChatResponse{}

	res, err := client.Chat(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "status_code [404] with body") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Chat_ErrorDoRequest(t *testing.T) {
	ctx := context.Background()
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("something went wrong")
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewChatRequest("llama3.2", []ollamaclient.Message{{Role: "user", Content: "why is the sky blue?"}}).
		WithStream(false).Build()
	expected := ollamaclient.ChatResponse{}

	res, err := client.Chat(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "do request:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}

func TestClient_Chat_ErrorCreatingRequest(t *testing.T) {
	ctx := context.Context(nil)
	c := ollamaclient.HTTPClientMock{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, errors.New("something went wrong")
		},
	}
	client := ollamaclient.NewClient(&c, "http://localhost:11434")
	req := ollamaclient.NewChatRequest("llama3.2", []ollamaclient.Message{{Role: "user", Content: "why is the sky blue?"}}).
		WithStream(false).Build()
	expected := ollamaclient.ChatResponse{}

	res, err := client.Chat(ctx, req)

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("expected [%+v] but was [%+v]", expected, res)
	}
	if err == nil {
		t.Fatal("an error was expected")
	}
	if !strings.HasPrefix(err.Error(), "creating request:") {
		t.Fatalf("unexpected message error: %v", err)
	}
}
