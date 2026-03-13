package services

import (
	"context"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/openweb3/web3go"
	"github.com/pkg/errors"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
	"github.com/wangdayong228/ydyl-console-service/internal/dtos"
	"github.com/wangdayong228/ydyl-console-service/internal/utils"
)

type ResultService struct {
	config configs.ResultFileConfig
}

type xjstBridgeInfoRaw struct {
	L1AdminAddress     string `json:"l1_admin_address"`
	L1SimpleCalculator string `json:"l1_simple_calculator_addr"`
	L1StateSender      string `json:"l1_state_sender_addr"`
	L1UnifiedBridge    string `json:"l1_unified_bridge_addr"`
	L2AdminAddress     string `json:"l2_admin_address"`
	L2SimpleCalculator string `json:"l2_simple_calculator_addr"`
	L2StateSender      string `json:"l2_state_sender_addr"`
	L2UnifiedBridge    string `json:"l2_unified_bridge_addr"`
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

	bridgInfoRaw, err := s.getXjstBridgeInfo()
	if err != nil {
		return nil, err
	}

	contracts.L1SimpleCalculator = common.HexToAddress(bridgInfoRaw.L1SimpleCalculator)
	contracts.L1StateSender = common.HexToAddress(bridgInfoRaw.L1StateSender)
	contracts.L1UnifiedBridge = common.HexToAddress(bridgInfoRaw.L1UnifiedBridge)
	contracts.L2StateSender = common.HexToAddress(bridgInfoRaw.L2StateSender)
	contracts.L2UnifiedBridge = common.HexToAddress(bridgInfoRaw.L2UnifiedBridge)

	return &contracts, nil
}

func (s *ResultService) getXjstBridgeInfo() (*xjstBridgeInfoRaw, error) {
	// 获取 l2 bridges : cast rpc layer2_getBridgeInfo --rpc-url http://44.233.198.160:30010 | jq
	client := web3go.MustNewClient("http://127.0.0.1:30010")

	bridgInfoRaw := &xjstBridgeInfoRaw{}

	err := client.Provider().CallContext(context.Background(), bridgInfoRaw, "layer2_getBridgeInfo")
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to get xjst bridge info")
	}
	return bridgInfoRaw, nil
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
		result.L1BridgeSendContract = contracts.L1StateSender
		result.L1BridgeReceiveContract = contracts.L1UnifiedBridge
		result.L2BridgeSendContract = contracts.L2StateSender
		result.L2BridgeReceiveContract = contracts.L2UnifiedBridge
		return result, nil

	case enums.L2TypeCdk:
		contracts, err := s.GetCdkNodeDeploymentContracts()
		if err != nil {
			return nil, err
		}
		result.L1BridgeSendContract = contracts.PolygonZkEVMBridgeAddress
		result.L1BridgeReceiveContract = contracts.PolygonZkEVMBridgeAddress
		result.L2BridgeSendContract = contracts.PolygonZkEVML2BridgeAddress
		result.L2BridgeReceiveContract = contracts.PolygonZkEVML2BridgeAddress
		return result, nil

	case enums.L2TypeOp:
		contracts, err := s.GetOpNodeDeploymentContracts()
		if err != nil {
			return nil, err
		}
		result.L1BridgeSendContract = contracts.L1CrossDomainMessengerProxy
		result.L1BridgeReceiveContract = contracts.L1CrossDomainMessengerProxy
		result.L2BridgeSendContract = contracts.L2CrossDomainMessenger
		result.L2BridgeReceiveContract = contracts.L2CrossDomainMessenger
		return result, nil
	default:
		return nil, errors.Errorf("unsupported L2_TYPE: %v", l2Type)
	}
}
