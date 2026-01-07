package utils

import (
	"errors"
	"os"

	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
)

func GetL2Type() (enums.L2Type, error) {
	l2TypeRaw, ok := os.LookupEnv("L2_TYPE")
	if !ok {
		return 0, errors.New("L2_TYPE is not set")
	}
	return enums.ParseL2Type(l2TypeRaw)
}
