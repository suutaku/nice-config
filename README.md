# nice-config
 a nice configure parser

## Use `nice-config` at your project

```go
import "github.com/suutaku/nice-config/pkg/conf"

func main(){
    // note: the_config_name must included a extension name. for now we supported json, yaml, toml by default.
    cnf := conf.NewConfigureParser("/the/config/file/home", "the_config_name.yaml")
    
     // use Load(interface) to load configure data and unmarshal to WhateverStruct.
    confStruct := WhateverStruct{}
    cnf.Load(&confStruct)

    // if you have some default values
    cnf.LoadWithMerge(&confStruct,toMergeWithSameStruct)

    // save your config data to file
    cnf.Save(confStruct)
}
```

## Define your own marshaller
we supported json, yaml, toml by default. if you want parse another formated files. just simply implement your mashaller as below. 

```go
type Marshaller interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
	Type() string
}
```
for example i want a json marshaller, firt implement it:

```go
type MyJson struct {
}

func NewMyJson() *MyJson {
	return &MyJson{}
}

func (js *MyJson) Type() string {
	return "myjson" // basically, this must be a file extension name
}

func (js *MyJson) Marshal(input interface{}) ([]byte, error) {
	return json.MarshalIndent(input, "", "\t")
}

func (js *MyJson) Unmarshal(input []byte, output interface{}) error {
	return json.Unmarshal(input, output)
}

```

sencond, regist it befor use:

```go

import "github.com/suutaku/nice-config/pkg/conf"

func main(){
  conf.RegistMarshaller(&MyJson{})
  /**
    TODO
  **/
}

```



