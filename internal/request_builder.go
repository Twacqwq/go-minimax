package minimax

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

type RequestBuilder interface {
	Build(ctx context.Context, url, method string, body any, header http.Header) (*http.Request, error)
}

type HTTPReqeustBuilder struct {
	marshaller Marshaller
}

func NewHTTPRequestBuilder() *HTTPReqeustBuilder {
	return &HTTPReqeustBuilder{
		marshaller: &JsonMarshaller{},
	}
}

func (builder *HTTPReqeustBuilder) Build(ctx context.Context, url, method string, body any, header http.Header) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		if v, ok := body.(io.Reader); ok {
			bodyReader = v
		} else {
			reqBytes, err := builder.marshaller.Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewBuffer(reqBytes)
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	if header != nil {
		req.Header = header
	}

	return req, nil
}
