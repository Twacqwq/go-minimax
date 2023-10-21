package minimax

import (
	"context"
	"net/http"
)

func (c *Client) CreateCompletion(ctx context.Context, request *ChatCompletionRequest, opts ...CompletionOption) (*ChatCompletionResponse, error) {
	if request.Stream {
		return nil, ErrCompletionStreamNotSupported
	}
	if !checkSupportModels(request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}

	initParam(request)
	for _, opt := range opts {
		opt(request)
	}

	req, err := c.newRequest(ctx, c.buildFullURL(request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &ChatCompletionResponse{}
	err = c.send(req, resp)

	return resp, err
}

func (c *Client) CreateCompletionStream(ctx context.Context, request *ChatCompletionRequest, opts ...CompletionOption) (*ChatCompletionStream, error) {
	request.Stream = true
	if !checkSupportModels(request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}

	initParam(request)
	for _, opt := range opts {
		opt(request)
	}

	req, err := c.newRequest(ctx, c.buildFullURL(request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendStream[ChatCompletionResponse](c, req)
	if err != nil {
		return nil, err
	}

	return &ChatCompletionStream{
		streamReader: resp,
	}, nil
}

func initParam(request *ChatCompletionRequest) {
	request.ReplyConstraints = ReplyConstraints{
		SenderType: ChatMessageRoleBot,
		SenderName: ModelBot,
	}
	request.BotSetting = []BotSetting{
		{
			BotName: ModelBot,
			Content: "MM智能助理是一款由MiniMax自研的, 没有调用其他产品的接口的大型语言模型。MiniMax是一家中国科技公司, 一直致力于进行大模型相关的研究.",
		},
	}
}

type CompletionOption func(*ChatCompletionRequest)

func WithReplyConstraints(v ReplyConstraints) CompletionOption {
	return func(cc *ChatCompletionRequest) {
		cc.ReplyConstraints = v
	}
}

func WithBotSetting(rolePrompt string, settings ...[]BotSetting) CompletionOption {
	return func(cc *ChatCompletionRequest) {
		cc.BotSetting[0].Content = rolePrompt
		for _, bot := range settings {
			cc.BotSetting = append(cc.BotSetting, bot...)
		}
	}
}
