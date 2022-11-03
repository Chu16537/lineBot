package conf

import (
	"io/ioutil"
	"linebot/dao"

	"gopkg.in/yaml.v2"
)

var Env *dao.Env

// 讀取 env檔案
func Init() error {
	envByte, err := ioutil.ReadFile("env.yaml")

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(envByte, &Env)

	if err != nil {
		return err
	}

	return nil
}
