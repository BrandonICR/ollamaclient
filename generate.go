package ollamaclient

type GenerateRequest struct {
	// Model (required) the model name.
	Model string `json:"model"`
	// Prompt the prompt to generate a response for.
	Prompt string `json:"prompt"`
	// Suffix the text after the model response.
	Suffix *string `json:"suffix"`
	// Images (optional) a list of base64-encoded images (for multimodal models such as llava).
	Images *string `json:"images"`
	// Format the format to return a response in. Format can be json or a JSON schema.
	Format *string `json:"format"`
	// Options additional model parameters listed in the documentation for the Modelfile such as temperature.
	Options *string `json:"options"`
	// System system message to (overrides what is defined in the Modelfile).
	System *string `json:"system"`
	// Template the prompt template to use (overrides what is defined in the Modelfile).
	Template *string `json:"template"`
	// Stream if false the response will be returned as a single response object, rather than a stream of objects.
	Stream *bool `json:"stream"`
	// Raw if true no formatting will be applied to the prompt. You may choose to use the raw parameter if you are specifying a full templated prompt in your request to the API.
	Raw *string `json:"raw"`
	// KeepAlive controls how long the model will stay loaded into memory following the request (default: 5m).
	KeepAlive *string `json:"keep_alive"`
}

type GenerateResponse struct {
	// Model (required) the model name.
	Model string `json:"model"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
	// Response empty if the response was streamed, if not streamed, this will contain the full response.
	Response string `json:"response"`
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

type GenerateRequestBuilder struct {
	request *GenerateRequest
}

func NewGenerateRequest(model, prompt string) *GenerateRequestBuilder {
	return &GenerateRequestBuilder{request: &GenerateRequest{
		Model:  model,
		Prompt: prompt,
	}}
}

func (g *GenerateRequestBuilder) WithSuffix(suffix string) *GenerateRequestBuilder {
	g.request.Suffix = &suffix
	return g
}

func (g *GenerateRequestBuilder) WithImages(images string) *GenerateRequestBuilder {
	g.request.Images = &images
	return g
}

func (g *GenerateRequestBuilder) WithFormat(format string) *GenerateRequestBuilder {
	g.request.Format = &format
	return g
}

func (g *GenerateRequestBuilder) WithOptions(options string) *GenerateRequestBuilder {
	g.request.Options = &options
	return g
}

func (g *GenerateRequestBuilder) WithSystem(system string) *GenerateRequestBuilder {
	g.request.System = &system
	return g
}

func (g *GenerateRequestBuilder) WithTemplate(template string) *GenerateRequestBuilder {
	g.request.Template = &template
	return g
}

func (g *GenerateRequestBuilder) WithStream(stream bool) *GenerateRequestBuilder {
	g.request.Stream = &stream
	return g
}

func (g *GenerateRequestBuilder) WithRaw(raw string) *GenerateRequestBuilder {
	g.request.Raw = &raw
	return g
}

func (g *GenerateRequestBuilder) WithKeepAlive(keepAlive string) *GenerateRequestBuilder {
	g.request.KeepAlive = &keepAlive
	return g
}

func (g *GenerateRequestBuilder) Build() GenerateRequest {
	return *g.request
}
