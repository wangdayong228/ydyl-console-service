package utils

import (
	"strconv"

	"github.com/joho/godotenv"
	"github.com/wangdayong228/ydyl-console-service/internal/constants/enums"
	"github.com/wangdayong228/ydyl-console-service/internal/dtos"
)

func ParsePipeProgress(filePath string) (*dtos.PipeProgressResponse, error) {
	entries, err := godotenv.Read(filePath)
	if err != nil {
		return nil, err
	}

	lastDoneStep, err := strconv.Atoi(entries["LAST_DONE_STEP"])
	if err != nil {
		return nil, err
	}

	return &dtos.PipeProgressResponse{
		LAST_DONE_STEP:  lastDoneStep,
		PIPELINE_STATUS: enums.PipelineStatus(entries["PIPELINE_STATUS"]),
	}, nil
}
