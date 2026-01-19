package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-goutils/utils/ginutils"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/services"
)

type ResultController struct {
	deployService *services.ResultService
}

func NewResultController() *ResultController {
	return &ResultController{
		deployService: services.NewResultService(configs.Get().ResultFile),
	}
}

// return deploy summary
func (c *ResultController) GetSummary(ctx *gin.Context) {
	result, err := c.deployService.GetSummary()
	ginutils.RenderResponse(ctx, result, err)
}

// return deploy pipeline progress
func (c *ResultController) GetPipelineProgress(ctx *gin.Context) {
	progress, err := c.deployService.GetPipeProgress()
	ginutils.RenderResponse(ctx, progress, err)
}

func (c *ResultController) GetNodeDeploymentContracts(ctx *gin.Context) {
	aliases, err := c.deployService.GetNodeDeploymentContracts()
	ginutils.RenderResponse(ctx, aliases, err)
}
