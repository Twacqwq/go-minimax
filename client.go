package minimax

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	pkg "github.com/Twacqwq/go-minimax/internal"
)

type Client struct {
	config         *Config
	requestBuilder *pkg.HTTPReqeustBuilder
}

func NewClient(apiToken, groupId string) *Client {
	return NewClientWithConfig(DefaultConfig(apiToken, groupId))
}

func NewClientWithConfig(config *Config) *Client {
	return &Client{
		config:         config,
		requestBuilder: pkg.NewHTTPRequestBuilder(),
	}
}

func (c *Client) newRequest(ctx context.Context, url, method string, opts ...option) (*http.Request, error) {
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}
	for _, opt := range opts {
		opt(args)
	}
	req, err := c.requestBuilder.Build(ctx, url, method, args.body, args.header)
	if err != nil {
		return nil, err
	}
	c.setCommonHeader(req)

	return req, nil
}

func (c *Client) setCommonHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.apiToken))
	req.Header.Set("Content-Type", "application/json")
}

func (c *Client) buildFullURL(version, model string) string {
	return fmt.Sprintf("%s%s?GroupId=%s", c.config.BaseURL, getURL(version, model), c.config.groupId)
}

func (c *Client) send(req *http.Request, v any, files ...*os.File) error {
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if isFailureStatusCode(resp) {
		return err
	}

	if resp.Header.Get("Content-Type") == "audio/mpeg" {
		return decodeResponse(resp.Body, v, files...)
	}

	return decodeResponse(resp.Body, v)
}

func sendStream[T steamable](client *Client, req *http.Request) (*streamReader[T], error) {
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.config.HTTPClient.Do(req)
	if err != nil || isFailureStatusCode(resp) {
		return new(streamReader[T]), err
	}

	return &streamReader[T]{
		emptyMessagesLimit: client.config.EmptyMessageLimit,
		reader:             bufio.NewReader(resp.Body),
		response:           resp,
		errAccumulator:     pkg.NewErrorAccumulator(),
		unmarshaler:        &pkg.JsonUnmarshaller{},
	}, nil
}

func decodeResponse(body io.Reader, v any, files ...*os.File) error {
	if v == nil {
		return nil
	}

	if res, ok := v.(*string); ok {
		return decodeString(body, res)
	}

	if len(files) > 0 {
		return decodeFile(body, files[0])
	}

	return json.NewDecoder(body).Decode(v)
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeFile(body io.Reader, output *os.File) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	_, err = output.Write(b)
	return err
}

func decodeString(body io.Reader, outout *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*outout = string(b)
	return nil
}

type option func(*requestOptions)

type requestOptions struct {
	body   any
	header http.Header
}

func withBody(body any) option {
	return func(opt *requestOptions) {
		opt.body = body
	}
}

// func withContentType(contentType string) option {
// 	return func(opt *requestOptions) {
// 		opt.header.Set("Content-Type", contentType)
// 	}
// }
