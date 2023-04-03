package api

import (
	"github.com/jinzhu/configor"
	"log"
	"os"
)

type ConfigBean struct {
	Parent   string
	Children []struct {
		Name  string
		Index int
	}
}

func ConfigorDemo() {
	//获取工程的跟目录
	log.Println(os.Getwd())
	jsonConfig := ConfigBean{}
	configor.Load(&jsonConfig, "./demo_project/hello_world/configs/config_json.json")
	log.Println(jsonConfig)

	yamlConfig := ConfigBean{}
	configor.Load(&yamlConfig, "./demo_project/hello_world/configs/config_yaml.yaml")
	log.Println(yamlConfig)
}
