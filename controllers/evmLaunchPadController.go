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

// @Summary Create New launchpad
// @Schemes
// @Description Create New launchpad In Evm
// @Tags EVM Launchpad
// @Produce json
// @Params tags body models.EvmLaunchpad true "Create New launchpad"
// @response default {object} models.EvmLaunchpaSuccessResponse{}
// @Router /evmLaunchpad/CreateNewLaunchPad [post]
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
	// 	],
	// 	"status": -1
	// }

	err := elc.EvmLaunchpadService.CreateNewLaunchpad(&req)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err)) // db 저장 실패 http 상태 코드
		return
	}


	response := models.EvmLaunchpaSuccessResponse{
		Status:  1,
		Message: "Create Launchpad Successful",
	}
	ctx.JSON(http.StatusOK, response)

}

type getMyAllLaunchpadReq struct {
	EoaAddress string `form:"eoa_address" binding:"startswith=0x"`
}

// @Summary getMy All EvmLaunchpad
// @Schemes
// @Description getMy All EvmLaunchpad
// @Tags EVM Launchpad
// @Produce json
// @Param eoa_address query string true "input my Eoa Address"
// @response default {object} models.EvmLaunchpaSuccessResponse{}
// @Router /evmLaunchpad/GetMyAllLaunchpad [get]
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

func (elc *EvmLaunchpadController) DeleteAllLaunchpadByAdmin(ctx *gin.Context) {
	// debug용 후에 데이터 전체 삭제해야 할 떄 사용
}

func (elc *EvmLaunchpadController) RegisterEvmLaunchpadRoutes(r *gin.Engine) {
	route := r.Group("/evmLaunchpad")

	route.POST("/CreateNewLaunchPad", elc.CreateNewLaunchPad)
	route.GET("/GetMyAllLaunchpad", elc.GetMyAllLaunchpad)
}
