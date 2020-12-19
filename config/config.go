package config

import (
	"flag"
	"fmt"
	"github.com/wb4646684/GoCommon/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configYAML struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Config *configYAML

func init() {
	ConfigFile := flag.String(
		"config",
		"config/config.yaml",
		"config file",
	)
	flag.Parse()
	Config = parseConfig(*ConfigFile)
	log.PrintLog(log.LogBody{
		Name:    "init",
		Message: fmt.Sprint("Load config: ", Config),
		Level:   log.INFO,
	})
}

func parseConfig(ConfigFile string) *configYAML {
	conf := &configYAML{}

	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return conf
}
