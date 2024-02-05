package minimax

import (
	"context"
	"fmt"
	"net/http"
)

type Thread struct {
	ID        string         `json:"id"`
	Object    string         `json:"object"`
	CreatedAt int64          `json:"created_at"`
	MetaData  map[string]any `json:"metadata,omitempty"`
	BaseResp  *BaseResp      `json:"base_resp"`
}

func (c *Client) CreateThreads(ctx context.Context, meta map[string]string) (*Thread, error) {
	req, err := c.newRequest(ctx, c.fullURL("/threads/create"), http.MethodPost, withBody(meta))
	if err != nil {
		return nil, err
	}

	resp := &Thread{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) RetrieveThreads(ctx context.Context, threadId string) (*Thread, error) {
	req, err := c.newRequest(ctx, fmt.Sprintf("%s&thread_id=%s", c.fullURL("/threads/retrieve"), threadId), http.MethodGet)
	if err != nil {
		return nil, err
	}

	resp := &Thread{}
	err = c.send(req, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
