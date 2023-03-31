package config

import (
	"go_study/logger"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type SettingConfig struct {
	JiraHostUrl string   `yaml:"JiraHostUrl"`
	Url         string   `yaml:"Url"`
	AssignUrl   string   `yaml:"AssignUrl"`
	UserName    string   `yaml:"UserName"`
	Password    string   `yaml:"Password"`
	TokenList   []string `yaml:"TokenList"`
}

var Cfg = &SettingConfig{}

func LoadConfig() {
	logger.Logger.Info("重新加载配置")
	var source = "./timer/cfg-dev.yml"
	if f, err := os.Open(source); err != nil {
		log.Fatalf("打开配置文件失败: %v", err)
	} else {
		if err := yaml.NewDecoder(f).Decode(Cfg); err != nil {
			log.Fatalf("反序列化配置文件失败: %v", err)
		}
	}
}
