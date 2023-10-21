package minimax

import "errors"

var (
	ErrCompletionUnsupportedModel   = errors.New("this model is not supported with this method, please use CreateChatCompletion client method instead") //nolint:lll
	ErrCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateCompletionStream")                      //nolint:lll
	ErrTooManyEmptyStreamMessages   = errors.New("")
)

var ()

const (
	Abab5     = "abab5-chat"
	Abab5Dot5 = "abab5.5-chat"
	Embo01    = "embo-01"

	ModelBot            = "MM智能助理"
	ChatMessageRoleUser = "USER"
	ChatMessageRoleBot  = "BOT"
	EmbeddingsDbType    = "db"
	EmbeddingsQueryType = "query"
)

var supportModels = map[string]string{
	Abab5:     "/text/chatcompletion",
	Abab5Dot5: "/text/chatcompletion_pro",
	Embo01:    "/embeddings",
}

func checkSupportModels(model string) bool {
	_, ok := supportModels[model]
	return ok
}

type Usage struct {
	TotalTokens           int64 `json:"total_tokens"`
	TokensWithAddedPlugin int64 `json:"tokens_with_added_plugin"`
}

type ChatCompletionRequest struct {
	Model             string           `json:"model"`
	Messages          []Message        `json:"messages"`
	BotSetting        []BotSetting     `json:"bot_setting"`
	SampleMessages    []Message        `json:"sample_messages,omitempty"`
	Stream            bool             `json:"stream,omitempty"`
	Prompt            string           `json:"prompt,omitempty"`
	TokensToGenerate  int64            `json:"tokens_to_generate,omitempty"`
	Temperature       float32          `json:"temperature,omitempty"`
	TopP              float32          `json:"top_p,omitempty"`
	MaskSensitiveInfo bool             `json:"mask_sensitive_info,omitempty"`
	Functions         []*Function      `json:"functions,omitempty"`
	FunctionCall      *FunctionCall    `json:"function_call,omitempty"`
	ReplyConstraints  ReplyConstraints `json:"reply_constraints"`
	Plugins           []string         `json:"plugins"`
}

type ChatCompletionResponse struct {
	ID              string              `json:"id"`
	Model           string              `json:"model"`
	Reply           string              `json:"reply"`
	Choices         []ChatMessageChoice `json:"choices"`
	Usage           Usage               `json:"usage"`
	InputSensitive  bool                `json:"input_sensitive,omitempty"`
	OutputSensitive bool                `json:"output_sensitive,omitempty"`
	BaseResp        BaseResp            `json:"base_resp,omitempty"`
}

type CreateEmbeddingsRequest struct {
	Model string `json:"model"`

	Texts []string `json:"texts"`
	Type  string   `json:"type"`
}

type CreateEmbeddingsResponse struct {
	Vectors  [][]float32 `json:"vectors"`
	BaseResp BaseResp    `json:"base_resp"`
}

type ChatCompletionStream struct {
	*streamReader[ChatCompletionResponse]
}

type ChatMessageChoice struct {
	FinishReason string    `json:"finish_reason,omitempty"`
	Messages     []Message `json:"messages"`
}

type Message struct {
	SenderType string `json:"sender_type"`
	SenderName string `json:"sender_name"`
	Text       string `json:"text"`
}

type BotSetting struct {
	BotName string `json:"bot_name"`
	Content string `json:"content"`
}

type BaseResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type ReplyConstraints struct {
	SenderType string `json:"sender_type"`
	SenderName string `json:"sender_name"`
	Glyph      *Glyph `json:"glyph,omitempty"`
}

type Glyph struct {
	Type           string `json:"type"`
	RawGlyph       string `json:"raw_glyph"`
	JsonProperties any    `json:"json_properties,omitempty"`
}

type Function struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Parameters  Parameters `json:"parameters"`
}

type FunctionCall struct {
	Type      string `json:"type,omitempty"`
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type Parameters struct {
	Type       string   `json:"type"`
	Required   []string `json:"required"`
	Properties any      `json:"properties"`
}
