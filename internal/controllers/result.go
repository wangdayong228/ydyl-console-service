package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nft-rainbow/rainbow-goutils/utils/ginutils"
	"github.com/wangdayong228/ydyl-console-service/internal/configs"
	"github.com/wangdayong228/ydyl-console-service/internal/dtos"
	"github.com/wangdayong228/ydyl-console-service/internal/services"
)

// 仅用于 swaggo 解析注释里的返回类型引用（避免 dtos 导入被 Go 判定为未使用）。
var _ = []any{
	(*dtos.SummaryResultResponse)(nil),
	(*dtos.PipeProgressResponse)(nil),
	(*dtos.GenAccSummaryResponse)(nil),
	(*dtos.NodeDeploymentContractsResponse)(nil),
	(*dtos.OpNodeDeploymentContracts)(nil),
	(*dtos.CdkNodeDeploymentContracts)(nil),
	(*dtos.XjstNodeDeploymentContracts)(nil),
}

type ResultController struct {
	deployService *services.ResultService
}

func NewResultController() *ResultController {
	return &ResultController{
		deployService: services.NewResultService(configs.Get().ResultFile),
	}
}

// @Tags			Result
// @ID				GetDeploySummary
// @Summary		获取部署汇总信息
// @Description	返回部署汇总信息（从部署结果文件读取）
// @Produce		json
// @Success		200	{object}	dtos.SummaryResultResponse
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/summary [get]
func (c *ResultController) GetSummary(ctx *gin.Context) {
	result, err := c.deployService.GetSummary()
	ginutils.RenderResponse(ctx, result, err)
}

// @Tags			Result
// @ID				GetDeployPipelineProgress
// @Summary		获取部署流水线进度
// @Description	返回部署流水线执行进度（从 pipeline state 文件解析）
// @Produce		json
// @Success		200	{object}	dtos.PipeProgressResponse
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/pipeline-progress [get]
func (c *ResultController) GetPipelineProgress(ctx *gin.Context) {
	progress, err := c.deployService.GetPipeProgress()
	ginutils.RenderResponse(ctx, progress, err)
}

// @Tags			Result
// @ID				GetGenAccSummary
// @Summary		获取批量生成账户汇总信息
// @Description	返回 ydyl-gen-accounts 聚合输出中的 summary 字段
// @Produce		json
// @Success		200	{object}	dtos.GenAccSummaryResponse
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/gen-acc/summary [get]
func (c *ResultController) GetGenAccSummary(ctx *gin.Context) {
	result, err := c.deployService.GetGenAccSummary()
	ginutils.RenderResponse(ctx, result, err)
}

// @Tags			Result
// @ID				GetNodeDeploymentContracts
// @Summary		获取节点部署合约地址（按 L2 类型聚合）
// @Description	根据当前 L2 类型（OP/CDK/XJST）读取相应部署合约文件，并返回关键桥合约地址
// @Produce		json
// @Success		200	{object}	dtos.NodeDeploymentContractsResponse
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/node-deployment-contracts [get]
func (c *ResultController) GetNodeDeploymentContracts(ctx *gin.Context) {
	aliases, err := c.deployService.GetNodeDeploymentContracts()
	ginutils.RenderResponse(ctx, aliases, err)
}

// @Tags			Result
// @ID				GetOpNodeDeploymentContracts
// @Summary		获取 OP 节点部署合约地址
// @Description	读取 OP 节点部署合约文件并返回合约地址
// @Produce		json
// @Success		200	{object}	dtos.OpNodeDeploymentContracts
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/node-deployment-contracts/op [get]
func (c *ResultController) GetOpNodeDeploymentContracts(ctx *gin.Context) {
	contracts, err := c.deployService.GetOpNodeDeploymentContracts()
	ginutils.RenderResponse(ctx, contracts, err)
}

// @Tags			Result
// @ID				GetCdkNodeDeploymentContracts
// @Summary		获取 CDK 节点部署合约地址
// @Description	读取 CDK 节点部署合约文件并返回合约地址
// @Produce		json
// @Success		200	{object}	dtos.CdkNodeDeploymentContracts
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/node-deployment-contracts/cdk [get]
func (c *ResultController) GetCdkNodeDeploymentContracts(ctx *gin.Context) {
	contracts, err := c.deployService.GetCdkNodeDeploymentContracts()
	ginutils.RenderResponse(ctx, contracts, err)
}

// @Tags			Result
// @ID				GetXjstNodeDeploymentContracts
// @Summary		获取 XJST 节点部署合约地址
// @Description	读取 XJST 节点部署合约文件并返回合约地址
// @Produce		json
// @Success		200	{object}	dtos.XjstNodeDeploymentContracts
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/node-deployment-contracts/xjst [get]
func (c *ResultController) GetXjstNodeDeploymentContracts(ctx *gin.Context) {
	contracts, err := c.deployService.GetXjstNodeDeploymentContracts()
	ginutils.RenderResponse(ctx, contracts, err)
}

// @Tags			Result
// @ID				GetXjstNodeDeploymentL1Contracts
// @Summary		获取 XJST 节点部署合约地址
// @Description	读取 XJST 节点部署合约文件并返回合约地址
// @Produce		json
// @Success		200	{object}	dtos.XjstNodeDeploymentL1Contracts
// @Failure		400	{object}	ginutils.GinErrorBody	"Bad request"
// @Failure		409	{object}	ginutils.GinErrorBody	"Conflict"
// @Failure		599	{object}	ginutils.GinErrorBody	"Business error"
// @Router			/v1/result/node-deployment-contracts/xjst/l1 [get]
func (c *ResultController) GetXjstNodeDeploymentL1Contracts(ctx *gin.Context) {
	contracts, err := c.deployService.GetXjstNodeDeploymentL1Contracts()
	ginutils.RenderResponse(ctx, contracts, err)
}
