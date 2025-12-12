package services

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/dtos"
)

type DeployService struct {
	config configs.DeployConfig
}

func NewDeployService(deployConfig configs.DeployConfig) *DeployService {
	return &DeployService{}
}

func (s *DeployService) GetDeployResult() (*dtos.DeployResultResponse, error) {
	content, err := os.ReadFile(s.config.ResultFile)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to read deploy result file")
	}
	var deployResult dtos.DeployResultResponse
	err = json.Unmarshal(content, &deployResult)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal deploy result")
	}
	return &deployResult, nil
}
