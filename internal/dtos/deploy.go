package dtos

import "github.com/wangdayong228/ydyl-console-service/internal/constants/enums"

type DeployResultResponse struct {
	L2Type                            enums.L2Type
	L2_RPC_URL                        string
	L1_VAULT_PRIVATE_KEY              string
	L2_VAULT_PRIVATE_KEY              string
	KURTOSIS_L1_PREALLOCATED_MNEMONIC string
	CLAIM_SERVICE_PRIVATE_KEY         string
	L2_PRIVATE_KEY                    string
	L1_CHAIN_ID                       string
	L2_CHAIN_ID                       string
	L1_RPC_URL                        string
	L2_COUNTER_CONTRACT               string
}
