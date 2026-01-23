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

func (s *ResultService) GetOpNodeDeploymentContracts() (*dtos.OpNodeDeploymentContracts, error) {
	contractFile, err := configs.Get().GetNodeDeploymentContractFile(enums.L2TypeOp)
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
	return &contracts, nil
}

func (s *ResultService) GetCdkNodeDeploymentContracts() (*dtos.CdkNodeDeploymentContracts, error) {
	contractFile, err := configs.Get().GetNodeDeploymentContractFile(enums.L2TypeCdk)
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
	return &contracts, nil
}

func (s *ResultService) GetXjstNodeDeploymentContracts() (*dtos.XjstNodeDeploymentContracts, error) {
	contractFile, err := configs.Get().GetNodeDeploymentContractFile(enums.L2TypeXjst)
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(contractFile)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to read xjst contract file: %s", contractFile)
	}

	var contracts dtos.XjstNodeDeploymentContracts
	if err := json.Unmarshal(content, &contracts); err != nil {
		return nil, errors.WithMessagef(err, "failed to unmarshal xjst contract file: %s", string(content))
	}
	return &contracts, nil
}

func (s *ResultService) GetNodeDeploymentContracts() (*dtos.NodeDeploymentContractsResponse, error) {
	l2Type, err := utils.GetL2Type()
	if err != nil {
		return nil, err
	}

	result := &dtos.NodeDeploymentContractsResponse{}

	switch l2Type {
	case enums.L2TypeXjst:
		contracts, err := s.GetXjstNodeDeploymentContracts()
		if err != nil {
			return nil, err
		}
		result.L2Bridge = contracts.L2StateSender
		result.L1Bridge = contracts.L1UnifiedBridge
		return result, nil

	case enums.L2TypeCdk:
		contracts, err := s.GetCdkNodeDeploymentContracts()
		if err != nil {
			return nil, err
		}
		result.L2Bridge = contracts.PolygonZkEVML2BridgeAddress
		result.L1Bridge = contracts.PolygonZkEVMBridgeAddress
		return result, nil

	case enums.L2TypeOp:
		contracts, err := s.GetOpNodeDeploymentContracts()
		if err != nil {
			return nil, err
		}
		result.L2Bridge = contracts.L2CrossDomainMessenger
		result.L1Bridge = contracts.L1CrossDomainMessengerProxy
		return result, nil
	default:
		return nil, errors.Errorf("unsupported L2_TYPE: %v", l2Type)
	}
}
