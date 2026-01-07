package enums

import (
	"fmt"
	"strconv"
	"strings"
)

// L2Type 表示部署的 L2 类型（与脚本/配置中的 L2_TYPE 对齐）
// - 兼容数字：0=cdk, 1=op, 2=xjst（脚本侧常用）
// - 兼容字符串：cdk/op/xjst（更可读）
type L2Type uint8

const (
	L2TypeCdk L2Type = iota
	L2TypeOp
	L2TypeXjst
)

func (t L2Type) String() string {
	switch t {
	case L2TypeCdk:
		return "cdk"
	case L2TypeOp:
		return "op"
	case L2TypeXjst:
		return "xjst"
	default:
		return fmt.Sprintf("unknown(%d)", t)
	}
}

func ParseL2Type(raw string) (L2Type, error) {
	s := strings.TrimSpace(strings.ToLower(raw))
	if s == "" {
		return 0, fmt.Errorf("invalid L2_TYPE: %s", raw)
	}

	// 先兼容数字输入（0/1/2）
	if n, err := strconv.Atoi(s); err == nil {
		switch L2Type(n) {
		case L2TypeCdk, L2TypeOp, L2TypeXjst:
			return L2Type(n), nil
		default:
			return 0, fmt.Errorf("invalid L2_TYPE: %s", raw)
		}
	}

	// 再兼容字符串输入（cdk/op/xjst）
	switch s {
	case "cdk":
		return L2TypeCdk, nil
	case "op":
		return L2TypeOp, nil
	case "xjst":
		return L2TypeXjst, nil
	default:
		return 0, fmt.Errorf("invalid L2_TYPE: %s", raw)
	}
}

func (t L2Type) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *L2Type) UnmarshalText(data []byte) error {
	val, err := ParseL2Type(string(data))
	if err != nil {
		return err
	}
	*t = val
	return nil
}
