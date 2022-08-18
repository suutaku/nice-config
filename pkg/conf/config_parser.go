package conf

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"
)

const (
	defaultBaseExt  = "yaml"
	defaultBaseName = "config"
)

type ConfigureParser struct {
	home  string
	name  string
	exten string
	m     Marshaller
}

func NewConfigureParser(home string, fileName string) *ConfigureParser {
	ret := &ConfigureParser{}
	absPath, err := filepath.Abs(home)
	if err != nil {
		panic(err)
	}
	ret.home = absPath
	_, err = os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(absPath, 0766)
			if err != nil {
				panic(err)
			}
		}
		panic(err)
	}
	ret.exten = filepath.Ext(fileName)
	ret.name = strings.TrimSuffix(fileName, ret.exten)
	if ret.exten == "" {
		ret.exten = defaultBaseExt
	}

	if ret.name == "" {
		ret.name = defaultBaseName
	}
	mName := strings.TrimPrefix(ret.exten, ".")
	ret.m = GetMarshaller(mName)
	if ret.m == nil {
		panic(fmt.Sprintf("unimplementation marshaller for extension %s", mName))
	}
	return ret
}

// load with default value
func (cnf *ConfigureParser) LoadWithMerge(output interface{}, toMerge interface{}) error {
	err := cnf.Load(output)
	if err != nil {
		return err
	}
	return mergo.Merge(output, toMerge)
}

// load config to output
func (cnf *ConfigureParser) Load(ouput interface{}) error {
	// read file
	fileBytes, err := os.ReadFile(path.Join(cnf.home, cnf.name+cnf.exten))
	if err != nil {
		if !os.IsExist(err) {
			return nil
		}
		return err
	}
	return cnf.m.Unmarshal(fileBytes, ouput)
}

// save input to config file
func (cnf *ConfigureParser) Save(input interface{}) error {
	f, err := os.OpenFile(path.Join(cnf.home, cnf.name+cnf.exten), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := cnf.m.Marshal(input)
	if err != nil {
		return err
	}
	f.Write(b)
	return nil
}
