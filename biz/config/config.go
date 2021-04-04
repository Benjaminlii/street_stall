package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var (
	AppConfig = BasicConfig{}
	once      = sync.Once{}
)

func InitConfig(configPath string) {
	once.Do(func() {
		loadConfig(configPath, &AppConfig)
	})
}

// loadConfig 从指定路径下的配置文件中读取配置信息，初始化到AppConfig中
func loadConfig(configPath string, baseConfig *BasicConfig) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("[system][config] config get err:%v", err)
	}
	err = yaml.Unmarshal(yamlFile, baseConfig)
	if err != nil {
		log.Fatalf("[system][config] unmarshal config yaml error:%v", err)
	}
	log.Print("[system][config] init config success!")
}
