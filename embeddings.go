package minimax

import (
	"context"
	"net/http"
)

func (c *Client) CreateEmbeddings(ctx context.Context, request *CreateEmbeddingsRequest) (*CreateEmbeddingsResponse, error) {
	request.Model = Embo01
	req, err := c.newRequest(ctx, c.buildFullURL(embedding, request.Model), http.MethodPost, withBody(request))
	if err != nil {
		return nil, err
	}
	resp := &CreateEmbeddingsResponse{}
	err = c.send(req, resp)

	return resp, err
}
