package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	_ "github.com/wangdayong228/ydyl-console-service/docs"
	"github.com/wangdayong228/ydyl-console-service/internal/controllers"
)

type PublicRouter struct {
	resultController *controllers.ResultController
}

func MustNewPublicRouter() *PublicRouter {
	return &PublicRouter{
		resultController: controllers.NewResultController(),
	}
}

func (p *PublicRouter) Setup(router *gin.Engine) {
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// 使用中间件的路由组
	v1 := router.Group("v1")

	result := v1.Group("result")
	{
		result.GET("/summary", p.resultController.GetSummary)
		result.GET("/pipeline-progress", p.resultController.GetPipelineProgress)
		result.GET("/node-deployment-contracts", p.resultController.GetNodeDeploymentContracts)
		result.GET("/node-deployment-contracts/op", p.resultController.GetOpNodeDeploymentContracts)
		result.GET("/node-deployment-contracts/cdk", p.resultController.GetCdkNodeDeploymentContracts)
		result.GET("/node-deployment-contracts/xjst", p.resultController.GetXjstNodeDeploymentContracts)
	}

	// user.Use(pkgmiddlewares.PaginationMiddleware)
	// {
	// 	user.POST("refresh-token", p.userController.RefreshToken)
	// 	user.GET("balance", p.userController.GetBalance)
	// 	user.GET("info", p.userController.GetInfo)
	// 	user.GET("txn/list", p.userController.GetTxnList)
	// 	user.GET("txn/info", p.userController.GetTxn)

	// 	kycNeedCheckPassportUsed := user.Group("kyc")
	// 	kycNeedCheckPassportUsed.Use(middlewares.PassportUsedCheck)
	// 	{
	// 		kycNeedCheckPassportUsed.POST("order", p.userController.CreateKycOrder)
	// 		kycNeedCheckPassportUsed.POST("order/reopen", p.userController.ReopenKycOrderIfFail)
	// 	}

	// 	kyc := user.Group("kyc")
	// 	{
	// 		kyc.POST("base", p.userController.SetKycBase)
	// 		kyc.POST("address", p.userController.SetKycAddress)
	// 		kyc.POST("photos", p.userController.SetKycPhotos)
	// 		kyc.POST("commitment", p.userController.SetKycCommitment)
	// 		// kyc.GET("order", p.userController.GetKycOrder)
	// 		kyc.GET("orders", p.userController.GetKycOrders)
	// 		kyc.GET("commiment-content", p.userController.GetKycCommitmentContent)

	// 		wasabi := kyc.Group("wasabi")
	// 		{
	// 			wasabi.GET("base", p.userController.GetWasabiBaseInfo)
	// 			wasabi.POST("base", p.userController.SetWasabiBaseInfo)
	// 			wasabi.GET("photos", p.userController.GetWasabiPhotosInfo)
	// 			wasabi.POST("photos", p.userController.SetWasabiPhotos)
	// 			wasabi.GET("completed", p.userController.GetWasabiCompletedState)
	// 		}
	// 	}

	// 	smsCode := user.Group("verification/sms-code")
	// 	{
	// 		smsCode.POST("send", p.userController.SendSmsCode)
	// 		// 非业务用接口，实际业务中在 kyc/base 中验证
	// 		smsCode.POST("verify", p.userController.VerifySmsCode)
	// 	}

	// 	withdraw := user.Group("orders/withdraw")
	// 	{
	// 		withdraw.POST("", p.userController.Withdraw)
	// 		withdraw.GET("", p.userController.GetWithdrawOrder)
	// 		withdraw.GET("/list", p.userController.GetUserWithdrawOrders)
	// 		withdraw.GET("/limit", p.userController.GetWithdrawLimit)
	// 	}

	// 	sbt := user.Group("orders/sbt")
	// 	{
	// 		sbt.POST("", p.userController.ApplySbt)
	// 		sbt.GET("latest", p.userController.GetUserLatestSbtOrder)
	// 	}

	// 	channel := user.Group("channel")
	// 	{
	// 		channel.POST("", p.userController.CreateChannel)
	// 		channel.GET("", p.userController.GetMyChannel)
	// 	}
	// }

	// card := v1.Group("card")
	// card.Use(middlewares.RecordRequest)
	// card.Use(pkgmiddlewares.PaginationMiddleware)
	// {
	// 	card.GET("fee/opencard", p.cardController.GetOpenCardFee)
	// 	orders := card.Group("orders")
	// 	{
	// 		open := orders.Group("open")
	// 		{
	// 			open.POST("", p.cardController.OpenCard)
	// 			open.GET("", p.cardController.GetOpenCardOrder)
	// 			open.GET("list", p.cardController.GetUserOpenCardOrders)
	// 		}

	// 		exchange := orders.Group("exchange")
	// 		{
	// 			exchange.GET("compliance", p.cardController.GetUserComplianceDetail)
	// 			exchange.POST("compliance", p.cardController.SubmitCompliance)
	// 			exchangeFiat := exchange.Group("fiat")
	// 			{
	// 				exchangeFiat.POST("", p.cardController.ExchangeFiat)
	// 				exchangeFiat.GET("", p.cardController.GetExchangeFiatOrder)
	// 				exchangeFiat.GET("list", p.cardController.GetUserExchangeFiatOrders)
	// 				exchangeFiat.GET("firstNonFail", p.cardController.GetFirstNonFailExchangeFiatOrder)
	// 			}
	// 		}

	// 	}

	// 	active := card.Group("active")
	// 	{
	// 		active.POST("", p.cardController.Active)
	// 	}
	// 	card.GET("list", p.cardController.ListCards)
	// 	card.GET("", p.cardController.GetCard)
	// 	card.GET("info", p.cardController.GetCardSenstiveInfo)
	// 	card.GET("limit", p.cardController.GetTodayExchangeRemainLimits)
	// 	card.GET("available", p.cardController.CheckAdminCardAndBalance)

	// 	card.POST("external", p.cardController.CreateExternalCard)
	// 	card.DELETE("external", p.cardController.DeleteExternalCard)
	// 	card.PUT("external", p.cardController.UpdateExternalCard)

	// 	user.GET("balance/:id", nil)
	// }

	// oracle := v1.Group("oracle")
	// {
	// 	oracle.GET("price", p.priceController.GetPrice)
	// }

	// file := v1.Group("files")
	// {
	// 	file.POST("kyc", p.fileController.UploadKycFile)
	// 	file.POST("wasabi", p.fileController.UploadWasabiFile)
	// 	file.POST("public", p.fileController.UploadPublicFile)
	// }
	// assets := v1.Group("assets")
	// {
	// 	assets.Use(middlewares.SelfPhotoAccessCheck).Static("/files/kyc", configs.Get().File.GetKycDir())
	// }
}
