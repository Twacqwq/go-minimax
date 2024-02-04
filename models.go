package minimax

import "errors"

var (
	ErrCompletionUnsupportedModel        = errors.New("this model is not supported with this method, please use CreateChatCompletion client method instead") //nolint:lll
	ErrCompletionStreamNotSupported      = errors.New("streaming is not supported with this method, please use CreateCompletionStream")                      //nolint:lll
	ErrTooManyEmptyStreamMessages        = errors.New("too many empty messages")
	ErrCreateTextToSpeechProNotSupported = errors.New("params is not supported, please use CreateTextToSpeechPro")
)

var (
	completion    = "/text/chatcompletion"
	completionPro = "/text/chatcompletion_pro"
	embedding     = "/embeddings"
	speech        = "/text_to_speech"
	speechPro     = "/t2a_pro"
)

const (
	Abab5       = "abab5-chat"
	Abab5Dot5   = "abab5.5-chat"
	Abab5Dot5s  = "abab5.5s-chat"
	Abab6       = "abab6-chat"
	Embo01      = "embo-01"
	Speech01    = "speech-01"
	Speech01Pro = "speech-01-pro" // alias speech-01=speech-01-pro

	ModelBot            = "MM智能助理"
	ChatMessageRoleUser = "USER"
	ChatMessageRoleBot  = "BOT"
	EmbeddingsDbType    = "db"
	EmbeddingsQueryType = "query"

	ToolCodeInterpreter = "code_interpreter"
	ToolRetrieval       = "retrieval"
	ToolFunction        = "function"
	ToolWebSearch       = "web_search"
)

// interface -> model:status
var supportModels = map[string]map[string]bool{
	completion: {
		Abab5:      false,
		Abab5Dot5:  true,
		Abab5Dot5s: true,
		Abab6:      false,
	},
	completionPro: {
		Abab5:      false,
		Abab5Dot5:  true,
		Abab5Dot5s: true,
		Abab6:      true,
	},
	embedding: {Embo01: true},
	speech:    {Speech01: true},
	speechPro: {Speech01Pro: true},
}

func checkSupportModels(version, model string) bool {
	return supportModels[version][model]
}

func getURL(version, module string) string {
	if supportModels[version][module] {
		return version
	}

	return ""
}

type Usage struct {
	TotalTokens           int64 `json:"total_tokens"`
	TokensWithAddedPlugin int64 `json:"tokens_with_added_plugin"`
}

type ChatCompletionRequest struct {
	Model               string    `json:"model"`
	Messages            []Message `json:"messages"`
	Stream              bool      `json:"stream,omitempty"`
	Prompt              string    `json:"prompt"`
	TokensToGenerate    int64     `json:"tokens_to_generate,omitempty"`
	Temperature         float32   `json:"temperature,omitempty"`
	TopP                float32   `json:"top_p,omitempty"`
	UseStandardSSE      bool      `json:"use_standard_sse,omitempty"`
	BeamWidth           int       `json:"beam_width,omitempty"`
	RoleMeta            *RoleMeta `json:"role_meta"`
	ContinueLastMessage bool      `json:"continue_last_message"`
	SkipInfoMask        bool      `json:"skip_info_mask"`
}

type ChatCompletionProRequest struct {
	Model             string           `json:"model"`
	Messages          []ProMessage     `json:"messages"`
	BotSetting        []BotSetting     `json:"bot_setting"`
	SampleMessages    []Message        `json:"sample_messages,omitempty"`
	Stream            bool             `json:"stream,omitempty"`
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
	ID                  string              `json:"id"`
	Created             int64               `json:"created"`
	Model               string              `json:"model"`
	Reply               string              `json:"reply"`
	Choices             []ChatMessageChoice `json:"choices"`
	Usage               Usage               `json:"usage"`
	InputSensitive      bool                `json:"input_sensitive,omitempty"`
	InputSensitiveType  int64               `json:"input_sensitive_type,omitempty"`
	OutputSensitive     bool                `json:"output_sensitive,omitempty"`
	OutputSensitiveType int64               `json:"output_sensitive_type"`
	BaseResp            BaseResp            `json:"base_resp,omitempty"`
}

type ChatCompletionProResponse struct {
	ID                  string                 `json:"id"`
	Created             int64                  `json:"created"`
	Model               string                 `json:"model"`
	Reply               string                 `json:"reply"`
	Choices             []ChatMessageProChoice `json:"choices"`
	Usage               Usage                  `json:"usage"`
	InputSensitive      bool                   `json:"input_sensitive,omitempty"`
	InputSensitiveType  int64                  `json:"input_sensitive_type,omitempty"`
	OutputSensitive     bool                   `json:"output_sensitive,omitempty"`
	OutputSensitiveType int64                  `json:"output_sensitive_type"`
	BaseResp            BaseResp               `json:"base_resp,omitempty"`
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

type ChatCompletionProStream struct {
	*streamReader[ChatCompletionProResponse]
}

type ChatMessageChoice struct {
	Text         string `json:"text"`
	Index        int64  `json:"index"`
	FinishReason string `json:"finish_reason,omitempty"`
	Delta        string `json:"delta"`
}

type ChatMessageProChoice struct {
	FinishReason string       `json:"finish_reason,omitempty"`
	Index        int64        `json:"index"`
	Messages     []ProMessage `json:"messages"`
}

type Message struct {
	SenderType string `json:"sender_type"`
	Text       string `json:"text"`
}

type ProMessage struct {
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

type TimberWeight struct {
	VoiceId string `json:"voice_id"`
	Weight  int    `json:"weight"`
}

type ExtraInfo struct {
	AudioLength     int64 `json:"audio_length,omitempty"`
	AudioSampleRate int64 `json:"audio_sample_rate,omitempty"`
	AudioSize       int64 `json:"audio_size,omitempty"`
	Bitrate         int64 `json:"bitrate,omitempty"`
	WordCount       int64 `json:"word_count,omitempty"`
}

type RoleMeta struct {
	UserName string `json:"user_name"`
	BotName  string `json:"bot_name"`
}

type Tool struct {
	Typ      string `json:"type,omitempty"`
	Function string `json:"function,omitempty"`
}
