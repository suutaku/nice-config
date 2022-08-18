package conf

import (
	"gopkg.in/yaml.v3"
)

type Yaml struct {
}

func NewYaml() *Yaml {
	return &Yaml{}
}

func (ym *Yaml) Type() string {
	return "yaml"
}

func (ym *Yaml) Marshal(input interface{}) ([]byte, error) {
	return yaml.Marshal(input)
}

func (ym *Yaml) Unmarshal(input []byte, output interface{}) error {
	return yaml.Unmarshal(input, output)
}
