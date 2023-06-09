package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Conf struct {
	HttpService server `yaml:"server"`
	Mysql       mysql  `yaml:"mysql"`
	Redis       redis  `yaml:"redis"`
	GinMode     string `yaml:"gin_mode"`
}

func (c *Conf) ConfInit() {
	//获取yaml路径
	wdPath, err := os.Getwd()
	if err != nil {
		log.Println("获取yaml路径错误：", err)
	}
	//fmt.Println(wdPath)
	ymlPath := wdPath + "/config/config-default.yml"

	// 读取 YAML 文件内容
	data, err := os.ReadFile(ymlPath)

	if err != nil {
		log.Println("读取yaml文件错误：", err)
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(data, c)
	if err != nil {
		log.Println("解析yaml数据错误：", err)
	}
	//fmt.Println(*Config)

}
