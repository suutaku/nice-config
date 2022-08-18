package conf

import "encoding/json"

type Json struct {
}

func NewJson() *Json {
	return &Json{}
}

func (js *Json) Type() string {
	return "json"
}

func (js *Json) Marshal(input interface{}) ([]byte, error) {
	return json.MarshalIndent(input, "", "\t")
}

func (js *Json) Unmarshal(input []byte, output interface{}) error {
	return json.Unmarshal(input, output)
}
