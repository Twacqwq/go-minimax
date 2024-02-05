package minimax

import (
	"context"
	"net/http"
)

type Run struct {
	ID             string            `json:"id"`
	Object         string            `json:"object"`
	CreatedAt      int64             `json:"created_at"`
	AssistantId    string            `json:"assistant_id"`
	ThreadId       string            `json:"thread_id"`
	Status         string            `json:"status"`
	StartedAt      int64             `json:"started_at"`
	ExpiresAt      int64             `json:"expires_at"`
	CancelledAt    int64             `json:"cancelled_at"`
	FailedAt       int64             `json:"failed_at"`
	CompletedAt    int64             `json:"completed_at"`
	LastError      *RunError         `json:"last_error"`
	Model          string            `json:"model"`
	Instructions   string            `json:"instructions"`
	Tools          []*RunTool        `json:"tools"`
	FileIds        []string          `json:"file_ids"`
	RequiredAction any               `json:"required_action"`
	Metadata       map[string]string `json:"metadata"`
}

type RunError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type RunTool struct {
	Typ      string       `json:"type"`
	Function *RunFunction `json:"function"`
}

type RunFunction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Parameters  any    `json:"parameters"`
}

type RunCreateRequest struct {
	ThreadId     string            `json:"thread_id"`
	AssistantId  string            `json:"assistant_id"`
	Model        string            `json:"model,omitempty"`
	Instructions string            `json:"instructions,omitempty"`
	Tools        []*RunTool        `json:"tools,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

type RequiredAction struct {
	Typ               string      `json:"type"`
	SubmitToolOutputs *ToolOutput `json:"submit_tool_outputs"`
}

type ToolOutput struct {
	ToolCallId string `json:"tool_call_id"`
	Output     string `json:"output"`
}

type RunResponse struct {
	*Run
	BaseResp *BaseResp `json:"base_resp"`
}

type RunRetrieveRequest struct {
	ThreadId string `json:"thread_id"`
	RunId    string `json:"run_id"`
}

type ListRunOption struct {
	ThreadId string `json:"thread_id"`
	Limit    int64  `json:"limit,omitempty"`
	Order    string `json:"order,omitempty"`
	After    string `json:"after,omitempty"`
	Before   string `json:"before,omitempty"`
}

type ListRunResponse struct {
	*RunResponse
	BaseResp *BaseResp `json:"base_resp"`
}

type SubmitRequest struct {
	ThreadId   string     `json:"thread_id"`
	RunId      string     `json:"run_id"`
	ToolOutput []*RunTool `json:"tool_output"`
}

func (c *Client) CreateRun(ctx context.Context, request *RunCreateRequest) (*RunResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/threads/run/create"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &RunResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveRun(ctx context.Context, request *RunRetrieveRequest) (*RunResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/threads/run/retrieve"), request), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &RunResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListRun(ctx context.Context, request *ListRunOption) (*ListRunResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/threads/run/list"), request), http.MethodGet, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &ListRunResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) SumbmitToolOutputsRun(ctx context.Context, request *SubmitRequest) (*RunResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/threads/run/submit_tool_outputs"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &RunResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
