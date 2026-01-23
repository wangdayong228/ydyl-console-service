package dtos

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
)

type SummaryResultResponse struct {
	L2Type                            enums.L2Type `json:",omitempty"`
	L2_RPC_URL                        string       `json:",omitempty"`
	L1_VAULT_PRIVATE_KEY              common.Hash  `json:",omitempty"`
	L2_VAULT_PRIVATE_KEY              common.Hash  `json:",omitempty"`
	KURTOSIS_L1_PREALLOCATED_MNEMONIC string       `json:",omitempty"`
	CLAIM_SERVICE_PRIVATE_KEY         common.Hash  `json:",omitempty"`
	// L2_PRIVATE_KEY 用途：
	// - L2 上部署 Counter 合约（bridge 注册流程依赖）
	// - ydyl-gen-accounts 的付款/部署账户（写入 ydyl-gen-accounts/.env 的 PRIVATE_KEY）
	L2_PRIVATE_KEY         common.Hash    `json:",omitempty"`
	L1_CHAIN_ID            string         `json:",omitempty"`
	L2_CHAIN_ID            string         `json:",omitempty"`
	L1_RPC_URL             string         `json:",omitempty"`
	L2_COUNTER_CONTRACT    common.Address `json:",omitempty"`
	L1_BRIDGE_HUB_CONTRACT common.Address `json:",omitempty"`
}

type NodeDeploymentContractsResponse struct {
	L2Bridge common.Address
	L1Bridge common.Address
}

type OpNodeDeploymentContracts struct {
	L2CrossDomainMessenger       common.Address // 0x4200000000000000000000000000000000000007
	L1CrossDomainMessengerProxy  common.Address
	L1StandardBridgeProxyAddress common.Address
	OptimismPortalProxy          common.Address
	DisputeGameFactoryProxy      common.Address
}

type CdkNodeDeploymentContracts struct {
	PolygonZkEVML2BridgeAddress common.Address
	PolygonZkEVMBridgeAddress   common.Address
}

type XjstNodeDeploymentContracts struct {
	L1SimpleCalculator common.Address
	L1StateSender      common.Address
	L1UnifiedBridge    common.Address
	L2StateSender      common.Address
	L2UnifiedBridge    common.Address
}

type PipeProgressResponse struct {
	LAST_DONE_STEP  int
	PIPELINE_STATUS enums.PipelineStatus
}
