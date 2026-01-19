package services

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
	"github.com/wangdayong228/ydyl-console-service/internal/dtos"
	"github.com/wangdayong228/ydyl-console-service/internal/utils"
)

type ResultService struct {
	config configs.ResultFileConfig
}

func NewResultService(resultFileConfig configs.ResultFileConfig) *ResultService {
	return &ResultService{
		config: resultFileConfig,
	}
}

func (s *ResultService) GetSummary() (*dtos.SummaryResultResponse, error) {
	resultFile, err := configs.Get().ResolveSummaryFilePath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(resultFile)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to read deploy result file: %s", resultFile)
	}
	var deployResult dtos.SummaryResultResponse
	err = json.Unmarshal(content, &deployResult)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to unmarshal deploy result: %s", string(content))
	}
	return &deployResult, nil
}

func (s *ResultService) GetPipeProgress() (*dtos.PipeProgressResponse, error) {
	stateFile, err := configs.Get().ResolvePipelineStateFilePath()
	if err != nil {
		return nil, err
	}

	return utils.ParsePipeProgress(stateFile)
}

func (s *ResultService) GetNodeDeploymentContracts() (*dtos.NodeDeploymentContractsResponse, error) {

	l2Type, err := utils.GetL2Type()
	if err != nil {
		return nil, err
	}

	deployResult, err := s.GetSummary()
	if err != nil {
		return nil, err
	}
	result := &dtos.NodeDeploymentContractsResponse{
		L2Type:      l2Type,
		L2Counter:   deployResult.L2_COUNTER_CONTRACT,
		L1BridgeHub: deployResult.L1_BRIDGE_HUB_CONTRACT,
	}

	switch l2Type {
	case enums.L2TypeXjst:
		return nil, errors.New("xjst 暂不支持获取合约别名")
	case enums.L2TypeCdk:
		contractFile, err := configs.Get().GetNodeDeploymentContractFile(l2Type)
		if err != nil {
			return nil, err
		}
		content, err := os.ReadFile(contractFile)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to read cdk contract file: %s", contractFile)
		}
		var contracts dtos.CdkNodeDeploymentContracts
		if err := json.Unmarshal(content, &contracts); err != nil {
			return nil, errors.WithMessagef(err, "failed to unmarshal cdk contract file: %s", string(content))
		}

		result.L2Bridge = contracts.PolygonZkEVML2BridgeAddress
		result.L1Bridge = contracts.PolygonZkEVMBridgeAddress
		return result, nil

	case enums.L2TypeOp:
		contractFile, err := configs.Get().GetNodeDeploymentContractFile(l2Type)
		if err != nil {
			return nil, err
		}
		content, err := os.ReadFile(contractFile)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to read op contract file: %s", contractFile)
		}

		var contracts dtos.OpNodeDeploymentContracts
		if err := json.Unmarshal(content, &contracts); err != nil {
			return nil, errors.WithMessagef(err, "failed to unmarshal op contract file: %s", string(content))
		}

		result.L2Bridge = contracts.L2CrossDomainMessenger
		result.L1Bridge = contracts.L1CrossDomainMessengerProxy
		return result, nil
	default:
		return nil, errors.Errorf("unsupported L2_TYPE: %v", l2Type)
	}
}
