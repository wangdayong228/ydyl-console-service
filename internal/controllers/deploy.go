package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-goutils/utils/ginutils"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/services"
)

type DeployController struct {
	deployService *services.DeployService
}

func NewDeployController() *DeployController {
	return &DeployController{
		deployService: services.NewDeployService(configs.Get().Deploy),
	}
}

// return deploy reesult
func (c *DeployController) GetResult(ctx *gin.Context) {
	result, err := c.deployService.GetDeployResult()
	ginutils.RenderResponse(ctx, result, err)
}

// return deploy pipeline progress
func (c *DeployController) GetPipelineProgress(ctx *gin.Context) {
	progress, err := c.deployService.GetPipeProgress()
	ginutils.RenderResponse(ctx, progress, err)
}

func (c *DeployController) GetContractAliases(ctx *gin.Context) {
	aliases, err := c.deployService.GetContractAliases()
	ginutils.RenderResponse(ctx, aliases, err)
}
