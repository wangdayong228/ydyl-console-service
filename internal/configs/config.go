package configs

import (
	"strings"

	"github.com/nft-rainbow/rainbow-goutils/utils/configutils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

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
	logrus.WithField("config", configVal).Info("config loaded")
}

func Get() *Config {
	return configVal
}

func (c *Config) CheckValid() error {
	if strings.TrimSpace(c.Deploy.ResultFile) == "" {
		return errors.Errorf("deploy result file is not valid: %s", c.Deploy.ResultFile)
	}
	return nil
}
