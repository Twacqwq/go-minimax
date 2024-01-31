package minimax

import (
	"context"
	"errors"
	"net/http"
)

var errInvalidParams = errors.New("this parameter is not supported, please use the Pro method")

func (c *Client) CreateCompletion(ctx context.Context, request *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	if request.Stream {
		return nil, ErrCompletionStreamNotSupported
	}
	if !checkSupportModels(completion, request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}
	if validateChat5Dot5Params(request) {
		return nil, errInvalidParams
	}

	req, err := c.newRequest(ctx, c.buildFullURL(completion, request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &ChatCompletionResponse{}
	err = c.send(req, resp)

	return resp, err
}

func (c *Client) CreateCompletionStream(ctx context.Context, request *ChatCompletionRequest) (*ChatCompletionStream, error) {
	request.Stream = true
	if !checkSupportModels(completion, request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}

	req, err := c.newRequest(ctx, c.buildFullURL(completion, request.Model), http.MethodPost, withBody(request))
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

func (c *Client) CreateCompletionPro(ctx context.Context, request *ChatCompletionProRequest, opts ...CompletionProOption) (*ChatCompletionProResponse, error) {
	if request.Stream {
		return nil, ErrCompletionStreamNotSupported
	}
	if !checkSupportModels(completionPro, request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}

	initParam(request)
	for _, opt := range opts {
		opt(request)
	}

	req, err := c.newRequest(ctx, c.buildFullURL(completionPro, request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &ChatCompletionProResponse{}
	err = c.send(req, resp)

	return resp, err
}

func (c *Client) CreateCompletionProStream(ctx context.Context, request *ChatCompletionProRequest, opts ...CompletionProOption) (*ChatCompletionProStream, error) {
	request.Stream = true
	if !checkSupportModels(completionPro, request.Model) {
		return nil, ErrCompletionUnsupportedModel
	}

	initParam(request)
	for _, opt := range opts {
		opt(request)
	}

	req, err := c.newRequest(ctx, c.buildFullURL(completionPro, request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp, err := sendStream[ChatCompletionProResponse](c, req)
	if err != nil {
		return nil, err
	}

	return &ChatCompletionProStream{
		streamReader: resp,
	}, nil
}

func initParam(request *ChatCompletionProRequest) {
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

type CompletionProOption func(*ChatCompletionProRequest)

func WithReplyConstraints(v ReplyConstraints) CompletionProOption {
	return func(cc *ChatCompletionProRequest) {
		cc.ReplyConstraints = v
	}
}

func WithBotSetting(rolePrompt string, settings ...[]BotSetting) CompletionProOption {
	return func(cc *ChatCompletionProRequest) {
		cc.BotSetting[0].Content = rolePrompt
		for _, bot := range settings {
			cc.BotSetting = append(cc.BotSetting, bot...)
		}
	}
}

func validateChat5Dot5Params(request *ChatCompletionRequest) bool {
	if request.Model != Abab5Dot5 {
		return false
	}

	if request.BeamWidth > 1 || request.ContinueLastMessage {
		return true
	}

	return false
}
