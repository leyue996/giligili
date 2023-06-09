package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Conf struct {
	HttpService server `yaml:"server"`
	Mysql       mysql  `yaml:"mysql"`
	Redis       redis  `yaml:"redis"`
}

func (c *Conf) ConfInit() error {
	//获取yaml路径
	wdPath, err := os.Getwd()
	if err != nil {
		return err
	}
	//fmt.Println(wdPath)
	ymlPath := wdPath + "/config/config-default.yml"

	// 读取 YAML 文件内容
	data, err := os.ReadFile(ymlPath)

	if err != nil {
		return err
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	//fmt.Println(*Config)
	return nil
}
