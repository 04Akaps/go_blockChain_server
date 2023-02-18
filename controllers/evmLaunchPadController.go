package controllers

import (
	"net/http"

	"go_blockChain_server/customerror"
	"go_blockChain_server/middleware"
	"go_blockChain_server/models"
	"go_blockChain_server/services"

	"github.com/gin-gonic/gin"
)

type EvmLaunchpadController struct {
	EvmLaunchpadService services.EvmLaunchpadService
}

func NewLaunchpadController(els services.EvmLaunchpadService) EvmLaunchpadController {
	return EvmLaunchpadController{
		EvmLaunchpadService: els,
	}
}

func (elc *EvmLaunchpadController) CreateNewLaunchPad(ctx *gin.Context) {
	var req models.EvmLaunchpad

	bodyCheckError := middleware.CheckBodyBinding(&req, ctx)
	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError, "status": -1})
		return
	}
	// error response example
	// {
	// 	"errors": [
	// 		{
	// 			"field": "EoaAddress",
	// 			"message": "This field is required"
	// 		},
	// 		{
	// 			"field": "ContractAddress",
	// 			"message": "This field is required"
	// 		},
	// 		{
	// 			"field": "NetworkChainId",
	// 			"message": "This field is required"
	// 		},
	// 		{
	// 			"field": "Price",
	// 			"message": "This field is required"
	// 		},
	// 		{
	// 			"field": "MetaDataUri",
	// 			"message": "This field is required"
	// 		}
	// 	],
	// 	"status": -1
	// }

	err := elc.EvmLaunchpadService.CreateNewLaunchpad(&req)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err)) // db 저장 실패 http 상태 코드
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 0, "message": "Create Launchpad"})
}

type getMyAllLaunchpadReq struct {
	EoaAddress string `form:"eoa_address" binding:"startswith=0x"`

	// startswith=0x
}

func (elc *EvmLaunchpadController) GetMyAllLaunchpad(ctx *gin.Context) {
	var req getMyAllLaunchpadReq

	queryCheckError := middleware.CheckUriQueryBinding(&req, ctx)
	if queryCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": queryCheckError, "status": -1})
		return
	}

	result, err := elc.EvmLaunchpadService.GetMyAllLaunchpad(req.EoaAddress)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (elc *EvmLaunchpadController) RegisterEvmLaunchpadRoutes(r *gin.Engine) {
	route := r.Group("/evmLaunchpad")

	route.POST("/CreateNewLaunchPad", elc.CreateNewLaunchPad)
	route.GET("/GetMyAllLaunchpad", elc.GetMyAllLaunchpad)
}
