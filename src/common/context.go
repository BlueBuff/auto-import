package common

import (
	"gopkg.in/yaml.v2"
	"hdg.com/db-demo/src/server/util"
)

var ConfigurationContext *ApplicationContext

func init() {
	context := new(ApplicationContext)
	err := context.Parse("resources/applicationContext.yaml")
	if err != nil {
		//兼容单元测试的问题
		err := context.Parse("../resources/applicationContext.yaml")
		if err != nil {
			panic(err)
		}
	}
	ConfigurationContext = context
}

type Configuration interface {
	Parse(path string) error
}

func (context *ApplicationContext) Parse(path string) error {
	reader := util.NewBufferFileReader(path)
	data, err := reader.Read()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, context)
	if err != nil {
		return err
	}
	return nil
}
