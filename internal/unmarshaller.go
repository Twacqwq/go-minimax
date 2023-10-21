package minimax

import "encoding/json"

type Unmarshaller interface {
	Unmarshal([]byte, any) error
}

type JsonUnmarshaller struct{}

func (ju *JsonUnmarshaller) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
