package config

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Type 配置文件类型定义
type Type struct {
	AppConfig AppType `yaml:"app"`
	DBConfig  DBType  `yaml:"db"`
	LogConfig LogType `yaml:"log"`
}

// GetAppConfig 获取 App 配置
func (c *Type) GetAppConfig() *AppType {
	return &c.AppConfig
}

// GetDBConfig 获取数据库配置
func (c *Type) GetDBConfig() *DBType {
	return &c.DBConfig
}

// GetLogConfig 获取日志配置
func (c *Type) GetLogConfig() *LogType {
	return &c.LogConfig
}

var cli = flag.String("mode", "pro", "Input app mode")

/*
NewConfig 获得配置
*/
func NewConfig() *Type {
	flag.Parse()
	log.Printf("[Init] 加载配置文件，运行模式为: %v", *cli)

	path := ""
	if *cli == "dev" {
		path = "application_dev.yml"
	} else {
		path = "application.yml"
	}

	var config Type
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicf("[Init] 打开配置文件失败: %v", err.Error())
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Panicf("[Init] 格式化配置文件失败: %v", err.Error())
	}
	config.AppConfig.Mode = *cli

	return &config
}
