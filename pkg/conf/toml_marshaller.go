package conf

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type Toml struct {
}

func NewToml() *Toml {
	return &Toml{}
}

func (tm *Toml) Type() string {
	return "toml"
}

func (tm *Toml) Marshal(input interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	err := toml.NewEncoder(&buf).Encode(input)
	return buf.Bytes(), err
}

func (tm *Toml) Unmarshal(input []byte, output interface{}) error {
	return toml.Unmarshal(input, output)
}
