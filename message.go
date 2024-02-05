package minimax

import (
	"context"
	"net/http"
)

type ListMessageOption struct {
	ThreadId string `json:"thread_id"`
	Limit    int64  `json:"limit,omitempty"`
	Order    string `json:"order,omitempty"`
	After    string `json:"after,omitempty"`
	Before   string `json:"before,omitempty"`
}

type AsstMessage struct {
	ID          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int64             `json:"created_at"`
	ThreadId    string            `json:"thread_id"`
	Role        string            `json:"role"`
	Content     []*Content        `json:"content"`
	FileIds     []string          `json:"file_ids"`
	AssistantId string            `json:"assistant_id"`
	RunId       string            `json:"run_id"`
	MetaData    map[string]string `json:"metadata"`
}

type Content struct {
	Typ       string       `json:"type"`
	Text      *TextContent `json:"text"`
	ImageFile *ImageFile   `json:"image_file"`
}

type TextContent struct {
	Value       string        `json:"value"`
	Annotations []*Annotation `json:"annotations"`
}

type ImageFile struct {
	FileId string `json:"file_id"`
}

type Annotation struct {
	Typ          string        `json:"type"`
	Text         string        `json:"text"`
	StartIndex   int64         `json:"start_index"`
	EndIndex     int64         `json:"end_index"`
	FileCitation *FileCitation `json:"file_citation"`
	WebCitation  *WebCitation  `json:"web_citation"`
}

type FileCitation struct {
	FileId string `json:"file_id"`
	Quote  string `json:"quote"`
}

type WebCitation struct {
	Url   string `json:"url"`
	Quote string `json:"quote"`
}

type ListMessagesResponse struct {
	Object   string         `json:"object"`
	Data     []*AsstMessage `json:"data"`
	FirstId  string         `json:"first_id"`
	LastId   string         `json:"last_id"`
	BaseResp *BaseResp      `json:"base_resp"`
}

type MessageCreateRequest struct {
	ThreadId string            `json:"thread_id"`
	Role     string            `json:"role"`
	Content  string            `json:"content"`
	FileIds  []string          `json:"file_ids,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

type MessageResponse struct {
	*AsstMessage
	BaseResp *BaseResp `json:"base_resp"`
}

type MessageRetrieveRequest struct {
	MessageId string `json:"message_id"`
	ThreadId  string `json:"thread_id"`
}

func (c *Client) ListMessages(ctx context.Context, request *ListMessageOption) (*ListMessagesResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/threads/messages/list"), request), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &ListMessagesResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateMessages(ctx context.Context, request *MessageCreateRequest) (*MessageResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/threads/messages/add"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &MessageResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveMessages(ctx context.Context, request *MessageRetrieveRequest) (*MessageResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/threads/messages/retrieve"), request), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &MessageResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
