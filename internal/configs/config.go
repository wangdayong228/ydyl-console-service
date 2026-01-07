package configs

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/nft-rainbow/rainbow-goutils/utils/configutils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wangdayong228/ydyl-console-service/internal/utils"
)

var (
	configVal                     *Config
	resultFilePathTemplate        *template.Template
	pipelineStateFilePathTemplate *template.Template
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Deploy DeployConfig `yaml:"deploy"`
}

type DeployConfig struct {
	ResultFile        string `yaml:"resultFile"`
	OpContractFile    string `yaml:"opContractFile"`
	CdkContractFile   string `yaml:"cdkContractFile"`
	PipelineStateFile string `yaml:"pipelineStateFile"`
}

func Init() {
	configVal = configutils.MustLoad[Config]()
	if err := configVal.CheckValid(); err != nil {
		panic(err)
	}
	logrus.WithField("config", configVal).Info("config loaded")

	resultFilePathTemplate = template.Must(template.New("resultFile").Parse(configVal.Deploy.ResultFile))
	pipelineStateFilePathTemplate = template.Must(template.New("pipelineState").Parse(configVal.Deploy.PipelineStateFile))
}

func Get() *Config {
	return configVal
}

func (c *Config) CheckValid() error {
	if strings.TrimSpace(c.Deploy.ResultFile) == "" {
		return errors.Errorf("deploy result file is not valid: %s", c.Deploy.ResultFile)
	}
	if strings.TrimSpace(c.Deploy.PipelineStateFile) == "" {
		return errors.Errorf("deploy pipeline state file is not valid: %s", c.Deploy.PipelineStateFile)
	}
	return nil
}

func (c *Config) ResolveDeployResultFilePath() (string, error) {
	l2Type, err := utils.GetL2Type()
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"L2Type": l2Type.String(),
	}
	var buf bytes.Buffer
	if err := resultFilePathTemplate.Execute(&buf, data); err != nil {
		return "", errors.WithMessage(err, "failed to execute resultFile template")
	}
	return buf.String(), nil
}

func (c *Config) ResolvePipelineStateFilePath() (string, error) {
	l2Type, err := utils.GetL2Type()
	if err != nil {
		return "", err
	}
	data := map[string]string{
		"L2Type": l2Type.String(),
	}
	var buf bytes.Buffer
	if err := pipelineStateFilePathTemplate.Execute(&buf, data); err != nil {
		return "", errors.WithMessage(err, "failed to execute pipelineStateFile template")
	}
	return buf.String(), nil
}
