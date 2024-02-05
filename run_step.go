package minimax

import (
	"context"
	"net/http"
)

type RunStep struct {
	ID          string            `json:"id"`
	Object      string            `json:"object"`
	CreatedAt   int64             `json:"created_at"`
	AssistantId string            `json:"assistant_id"`
	ThreadId    string            `json:"thread_id"`
	Typ         string            `json:"type"`
	Status      string            `json:"status"`
	StartedAt   int64             `json:"started_at"`
	ExpiresAt   int64             `json:"expires_at"`
	CancelledAt int64             `json:"cancelled_at"`
	FailedAt    int64             `json:"failed_at"`
	CompletedAt int64             `json:"completed_at"`
	LastError   map[string]string `json:"last_error"`
	StepDetails *StepDetails      `json:"step_details"`
}

type StepDetails struct {
	Typ             string           `json:"type"`
	MessageCreation *MessageCreation `json:"message_creation"`
	ToolCalls       []*ToolCall      `json:"tool_calls"`
}

type MessageCreation struct {
	MessageId string `json:"message_id"`
}

type ToolCall struct {
	Typ             string           `json:"type"`
	CodeInterpreter *CodeInterpreter `json:"code_interpreter"`
	WebSearch       *WebSearch       `json:"web_search"`
	Retrieval       *CallRetrieval   `json:"retrieval"`
	Function        *CallFunction    `json:"function"`
}

type CodeInterpreter struct {
	ID              string `json:"id"`
	Typ             string `json:"type"`
	CodeInterpreter struct {
		Input   string `json:"input"`
		Outputs any    `json:"outputs"`
	} `json:"code_interpreter"`
}

type WebSearch struct {
	ID        string `json:"id"`
	Typ       string `json:"type"`
	WebSearch struct {
		Query   string `json:"query"`
		Outputs string `json:"outputs"`
	} `json:"web_search"`
}

type CallRetrieval struct {
	ID        string `json:"id"`
	Typ       string `json:"type"`
	Retrieval struct {
		Query   string `json:"query"`
		Outputs string `json:"outputs"`
	} `json:"retrieval"`
}

type CallFunction struct {
	ID       string `json:"id"`
	Typ      string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
		Output    string `json:"output"`
	} `json:"function"`
}

type RunStepRetrieveRequest struct {
	ThreadId string `json:"thread_id"`
	RunId    string `json:"run_id"`
	StepId   string `json:"step_id"`
}

type RunStepResponse struct {
	*RunStep
	BaseResp *BaseResp `json:"base_resp"`
}

type RunStepOption struct {
	ThreadId string `json:"thread_id"`
	RunId    string `json:"run_id"`
	Limit    int64  `json:"limit,omitempty"`
	Order    string `json:"order,omitempty"`
	After    string `json:"after,omitempty"`
	Before   string `json:"before,omitempty"`
}

type RunStepListResponse struct {
	Object   string             `json:"object"`
	Data     []*RunStepResponse `json:"data"`
	BaseResp *BaseResp          `json:"base_resp"`
}

func (c *Client) RetrieveRunStep(ctx context.Context, request *RunStepRetrieveRequest) (*RunStepResponse, error) {
	req, err := c.newRequest(ctx, c.fullURL("/threads/run_steps/retrieve"), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}

	resp := &RunStepResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) ListRunStep(ctx context.Context, opt *RunStepOption) (*RunStepListResponse, error) {
	req, err := c.newRequest(ctx, c.queryBuilder.Build(c.fullURL("/threads/run_steps/list"), opt), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &RunStepListResponse{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
