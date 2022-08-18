package conf

type Marshaller interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	Type() string
}

var enabledMarshaller = []Marshaller{
	&Yaml{},
	&Json{},
	&Toml{},
}

// get marshaller implementation by type name
func GetMarshaller(name string) Marshaller {
	for i, v := range enabledMarshaller {
		if v.Type() == name {
			return enabledMarshaller[i]
		}
	}
	return nil
}

// registe your marshaller implementation
func RegistMarshaller(input Marshaller) {
	enabledMarshaller = append(enabledMarshaller, input)
}
