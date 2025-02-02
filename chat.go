package ollamaclient

import "strings"

type ChatRequest struct {
	// Model (required) the model name.
	Model string `json:"model"`
	// Messages the messages of the chat, this can be used to keep a chat memory.
	Messages []Message `json:"messages"`
	// Tools list of tools in JSON for the model to use if supported.
	Tools *string `json:"tools"`
	// Format the format to return a response in. Format can be json or a JSON schema.
	Format *string `json:"format"`
	// Options additional model parameters listed in the documentation for the Modelfile such as temperature.
	Options *string `json:"options"`
	// Stream if false the response will be returned as a single response object, rather than a stream of objects.
	Stream *bool `json:"stream"`
	// KeepAlive controls how long the model will stay loaded into memory following the request (default: 5m).
	KeepAlive *string `json:"keep_alive"`
}

type ChatResponse struct {
	// Model (required) the model name.
	Model string `json:"model"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
	// Message
	Message Message `json:"message"`
	// Done
	Done bool `json:"done"`
	// DoneReason
	DoneReason string `json:"done_reason"`
	// Context an encoding of the conversation used in this response, this can be sent in the next request to keep a conversational memory.
	Context []int `json:"context"`
	// TotalDuration time spent generating the response.
	TotalDuration int `json:"total_duration"`
	// LoadDuration time spent in nanoseconds loading the model.
	LoadDuration int `json:"load_duration"`
	// PromptEvalCount number of tokens in the prompt.
	PromptEvalCount int `json:"prompt_eval_count"`
	// PromptEvalDuration time spent in nanoseconds evaluating the prompt.
	PromptEvalDuration int `json:"prompt_eval_duration"`
	// EvalCount number of tokens in the response.
	EvalCount int `json:"eval_count"`
	// EvalDuration time in nanoseconds spent generating the response.
	EvalDuration int `json:"eval_duration"`
}

type Message struct {
	// Role the role of the message, either system, user, assistant, or tool
	Role string `json:"role"`
	// Content the content of the message
	Content string `json:"content"`
	// Images a list of images to include in the message (for multimodal models such as llava)
	Images *string `json:"images"`
	// ToolCalls a list of tools in JSON that the model wants to use
	ToolCalls *string `json:"tool_calls"`
}

func CutThink(content string) string {
	index := strings.Index(content, "</think>")
	if index == -1 {
		return content
	}
	return content[index+8:]
}

type ChatRequestBuilder struct {
	request *ChatRequest
}

func NewChatRequest(model string, messages []Message) *ChatRequestBuilder {
	return &ChatRequestBuilder{request: &ChatRequest{
		Model:    model,
		Messages: messages,
	}}
}

func (g *ChatRequestBuilder) WithFormat(format string) *ChatRequestBuilder {
	g.request.Format = &format
	return g
}

func (g *ChatRequestBuilder) WithTools(tools string) *ChatRequestBuilder {
	g.request.Tools = &tools
	return g
}

func (g *ChatRequestBuilder) WithOptions(options string) *ChatRequestBuilder {
	g.request.Options = &options
	return g
}

func (g *ChatRequestBuilder) WithStream(stream bool) *ChatRequestBuilder {
	g.request.Stream = &stream
	return g
}

func (g *ChatRequestBuilder) WithKeepAlive(keepAlive string) *ChatRequestBuilder {
	g.request.KeepAlive = &keepAlive
	return g
}

func (g *ChatRequestBuilder) Build() ChatRequest {
	return *g.request
}
