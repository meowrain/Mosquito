package conf

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type GlobalConfiguration struct {
	App AppConfiguration `yaml:"app"`
}
type AppConfiguration struct {
	Name           string `yaml:"name"`
	Version        string `yaml:"version"`
	Host           string `yaml:"host"`
	Port           uint   `yaml:"port"`
	MaxConnections int    `yaml:"maxConnections"`
	MaxPackageSize int    `yaml:"maxPackageSize"`
}

var GlobalConf GlobalConfiguration

func (g *GlobalConfiguration) Load() error {
	confContent, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(confContent, &GlobalConf)
	if err != nil {
		return err
	}
	return nil
}
func init() {
	GlobalConf = GlobalConfiguration{
		App: AppConfiguration{
			Name:           "Mosquito",
			Version:        "0.2",
			Host:           "0.0.0.0",
			Port:           8099,
			MaxConnections: 1000,
			MaxPackageSize: 4096,
		},
	}
	err := GlobalConf.Load()
	if err != nil {
		log.Println("config file read failed:", err)
	}
}
