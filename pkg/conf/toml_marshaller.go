package conf

import (
	"github.com/naoina/toml"
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
	return toml.Marshal(input)
}

func (tm *Toml) Unmarshal(input []byte, output interface{}) error {
	return toml.Unmarshal(input, output)
}
