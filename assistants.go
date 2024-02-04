package minimax

import (
	"context"
	"fmt"
	"net/http"
)

type Assistant struct {
	ID           string            `json:"id"`
	Object       string            `json:"object"`
	CreatedAt    int64             `json:"created_at"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Model        string            `json:"model"`
	Instructions string            `json:"instructions"`
	Tools        []*Tool           `json:"tools"`
	FileIds      []string          `json:"file_ids"`
	MetaData     map[string]string `json:"meta_data"`
	RoleMeta     *RoleMeta         `json:"role_meta"`
	Status       string            `json:"status"`
}

type AssistantListOption struct {
	AssistantId string `json:"assistant_id"`
	FileId      string `json:"file_id"`
	Limit       int    `json:"limit"`
	Order       string `json:"order"`
	Before      string `json:"before"`
	After       string `json:"after"`
}

type AssistantListResponse struct {
	Object   string       `json:"object"`
	Data     []*Assistant `json:"data"`
	HasMore  bool         `json:"has_more"`
	FirstId  string       `json:"first_id"`
	LastId   string       `json:"last_id"`
	BaseResp *BaseResp    `json:"base_resp"`
}

type AssistantCreateRequest struct {
	Model        string            `json:"model"`
	RoleMeta     *RoleMeta         `json:"role_meta"`
	Instructions string            `json:"instructions,omitempty"`
	Name         string            `json:"name,omitempty"`
	Description  string            `json:"description,omitempty"`
	Tools        []*Tool           `json:"tools,omitempty"`
	FileIds      []string          `json:"file_ids,omitempty"`
	MetaData     map[string]string `json:"meta_data,omitempty"`
}

type AssistantRetrieveResponse struct {
	*Assistant
	BaseResp *BaseResp `json:"base_resp"`
}

type AssistantDeleteResponse struct {
	ID       string    `json:"id"`
	Object   string    `json:"object"`
	Deleted  bool      `json:"deleted"`
	BaseResp *BaseResp `json:"base_resp"`
}

type AssistantFile struct {
	ID          string `json:"id"`
	Object      string `json:"object"`
	CreatedAt   int64  `json:"created_at"`
	AssistantId string `json:"assistant_id"`
}

type AssistantFileListResponse struct {
	Object   string           `json:"object"`
	Data     []*AssistantFile `json:"data"`
	HasMore  bool             `json:"has_more"`
	FirstId  string           `json:"first_id"`
	LastId   string           `json:"last_id"`
	BaseResp *BaseResp        `json:"base_resp"`
}

type AssistantFilesCreateRequest struct {
	AssistantId string `json:"assistant_id"`
	FileId      string `json:"file_id"`
}

type AssistantFilesCreateResponse struct {
	ID          string    `json:"id"`
	Object      string    `json:"object"`
	CreatedAt   string    `json:"created_at"`
	AssistantId string    `json:"assistant_id"`
	BaseResp    *BaseResp `json:"base_resp"`
}

func (c *Client) ListAssistants(ctx context.Context, opt *AssistantListOption) (*AssistantListResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/assistants/list"), opt), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &AssistantListResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateAssistants(ctx context.Context, request *AssistantCreateRequest) (*Assistant, error) {
	req, err := c.newRequest(ctx, c.fullURL("/assistants/create"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &Assistant{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveAssistants(ctx context.Context, assistantId string) (*AssistantRetrieveResponse, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("%s&assistant_id=%s", c.fullURL("/assistants/retrieve"), assistantId), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &AssistantRetrieveResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) DeleteAssistants(ctx context.Context, assistantId string) (*AssistantDeleteResponse, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("%s&assistant_id=%s", c.fullURL("/assistants/delete"), assistantId), http.MethodPost)
	if err != nil {
		return nil, err
	}

	resp := &AssistantDeleteResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListAssistantFiles(ctx context.Context, opt *AssistantListOption) (*AssistantFileListResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/assistants/files/list"), opt), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &AssistantFileListResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) CreateAssistantFiles(ctx context.Context, request *AssistantFilesCreateRequest) (*AssistantFilesCreateResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/assistants/files/create"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &AssistantFilesCreateResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveAssistantFiles(ctx context.Context, opt *AssistantListOption) (*AssistantFilesCreateResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/assistants/files/retrieve"), opt), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &AssistantFilesCreateResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
