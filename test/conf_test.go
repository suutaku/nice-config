package test

import (
	"reflect"
	"testing"

	"github.com/suutaku/nice-config/pkg/conf"
)

type testConf struct {
	Name      string
	Port      int32
	Address   []string
	BadPerson bool
}

var defaultTestConf = testConf{
	Name:    "test",
	Port:    1024,
	Address: []string{"sichuan", "chengdu", "tianfu"},
}

func Test_Load(t *testing.T) {
	tests := []struct {
		name string
		args testConf
		want testConf
	}{
		{"yaml", defaultTestConf, defaultTestConf},
		{"json", defaultTestConf, defaultTestConf},
		{"toml", defaultTestConf, defaultTestConf},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Load(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Load(typeName string) testConf {
	cnf := conf.NewConfigureParser("/tmp", "load."+typeName)
	ret := testConf{}
	err := cnf.LoadWithMerge(&ret, defaultTestConf)
	if err != nil {
		panic(err)
	}
	return ret
}
