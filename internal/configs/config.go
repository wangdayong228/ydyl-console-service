package configs

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/nft-rainbow/rainbow-goutils/utils/configutils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
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

	ResultFile ResultFileConfig `yaml:"resultFile"`
}

type ResultFileConfig struct {
	Summary                    string                  `yaml:"summary"`
	NodeDeploymentContracts    NodeDeploymentContracts `yaml:"nodeDeploymentContracts"`
	L2counterAndRegisterBridge string                  `yaml:"l2counterAndRegisterBridge"`
	PipelineState              string                  `yaml:"pipelineState"`
}

type NodeDeploymentContracts struct {
	Op   string `yaml:"op"`
	Cdk  string `yaml:"cdk"`
	Xjst string `yaml:"xjst"`
}

func Init() {
	configVal = configutils.MustLoad[Config]()
	if err := configVal.CheckValid(); err != nil {
		panic(err)
	}
	logrus.WithField("config", configVal).Info("config loaded")

	resultFilePathTemplate = template.Must(template.New("resultFile").Parse(configVal.ResultFile.Summary))
	pipelineStateFilePathTemplate = template.Must(template.New("pipelineState").Parse(configVal.ResultFile.PipelineState))
}

func Get() *Config {
	return configVal
}

func (c *Config) CheckValid() error {
	if strings.TrimSpace(c.ResultFile.Summary) == "" {
		return errors.Errorf("result file summary is not valid: %s", c.ResultFile.Summary)
	}
	if strings.TrimSpace(c.ResultFile.PipelineState) == "" {
		return errors.Errorf("result file pipeline state is not valid: %s", c.ResultFile.PipelineState)
	}
	return nil
}

func (c *Config) ResolveSummaryFilePath() (string, error) {
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

func (c *Config) GetNodeDeploymentContractFile(l2Type enums.L2Type) (string, error) {
	switch l2Type {
	case enums.L2TypeOp:
		contractFile := strings.TrimSpace(c.ResultFile.NodeDeploymentContracts.Op)
		if contractFile == "" {
			return "", errors.New("op contract file is not set in config")
		}
		return contractFile, nil
	case enums.L2TypeCdk:
		contractFile := strings.TrimSpace(c.ResultFile.NodeDeploymentContracts.Cdk)
		if contractFile == "" {
			return "", errors.New("cdk contract file is not set in config")
		}
		return contractFile, nil
	case enums.L2TypeXjst:
		contractFile := strings.TrimSpace(c.ResultFile.NodeDeploymentContracts.Xjst)
		if contractFile == "" {
			return "", errors.New("xjst contract file is not set in config")
		}
		return contractFile, nil
	default:
		return "", errors.Errorf("unsupported L2_TYPE: %v", l2Type)
	}
}
