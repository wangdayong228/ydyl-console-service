package configs

import "github.com/nft-rainbow/rainbow-goutils/utils/configutils"

var configVal *Config

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Deploy DeployConfig `yaml:"deploy"`
}

type DeployConfig struct {
	ResultFile string `yaml:"resultFile"`
}

func Init() {
	configVal = configutils.MustLoad[Config]()
	if err := configVal.CheckValid(); err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &Config{}
}

func (c *Config) CheckValid() error {
	return nil
}
