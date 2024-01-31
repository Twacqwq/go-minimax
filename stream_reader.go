package minimax

import (
	"bufio"
	"bytes"
	"io"
	"net/http"

	pkg "github.com/Twacqwq/go-minimax/internal"
)

var (
	headerData  = []byte("data: ")
	errorPrefix = []byte(`data: {"error":`)
)

type steamable interface {
	ChatCompletionResponse | ChatCompletionProResponse
}

type streamReader[T steamable] struct {
	emptyMessagesLimit uint
	isFinished         bool

	reader         *bufio.Reader
	response       *http.Response
	errAccumulator pkg.ErrorAccumulator
	unmarshaler    pkg.Unmarshaller
}

func (stream *streamReader[T]) Recv() (T, error) {
	if stream.isFinished {
		return *new(T), io.EOF
	}

	return stream.process()
}

func (stream *streamReader[T]) process() (T, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
	)

	for {
		rawLine, readErr := stream.reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			return *new(T), readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.errAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return *new(T), writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.emptyMessagesLimit {
				return *new(T), ErrTooManyEmptyStreamMessages
			}

			continue
		}

		var response T
		unmarshalErr := stream.unmarshaler.Unmarshal(bytes.TrimPrefix(noSpaceLine, headerData), &response)
		if unmarshalErr != nil {
			return *new(T), unmarshalErr
		}

		return response, nil
	}
}

func (stream *streamReader[T]) Close() {
	stream.response.Body.Close()
}
