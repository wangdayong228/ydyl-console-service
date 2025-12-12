package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-goutils/utils/ginutils"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/services"
)

type DeployController struct {
}

func NewDeployController() *DeployController {
	return &DeployController{}
}

// return deploy reesult
func (c *DeployController) GetResult(ctx *gin.Context) {
	result, err := services.NewDeployService(configs.Get().Deploy).GetDeployResult()
	ginutils.RenderResponse(ctx, result, err)
}
