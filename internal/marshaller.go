package minimax

import "encoding/json"

type Marshaller interface {
	Marshal(any) ([]byte, error)
}

type JsonMarshaller struct{}

func (jm *JsonMarshaller) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
