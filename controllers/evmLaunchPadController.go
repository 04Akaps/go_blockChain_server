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

	bodyCheckError := middleware.CheckBodyBinding(req, ctx)
	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError})
	}

	// error response example
	// {
	// 	"errors": [
	// 		{
	// 			"field": "Name",
	// 			"message": "This field is required"
	// 		}
	// 	]
	// }

	err := elc.EvmLaunchpadService.CreateNewLaunchpad(&req)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err)) // db 저장 실패 http 상태 코드
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 0, "message": "Create Launchpad"})
}

type getMyAllLaunchpadReq struct {
	EoaAddress string `json:"eoa_address" binding:"required,startswith=0x"`
}

func (elc *EvmLaunchpadController) GetMyAllLaunchpad(ctx *gin.Context) {
	var req getMyAllLaunchpadReq

	bodyCheckError := middleware.CheckBodyBinding(req, ctx)
	if bodyCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": bodyCheckError})
	}

	// if err := ctx.ShouldBindJSON(&req); err != nil {
	// 	// bind 체크를 위한 코드
	// 	var ve validator.ValidationErrors
	// 	if errors.As(err, &ve) {
	// 		out := make([]models.ErrorMsg, len(ve))
	// 		for i, fe := range ve {
	// 			out[i] = models.ErrorMsg{fe.Field(), customerror.GetErrorMsg(fe)}
	// 		}
	// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
	// 	}
	// 	return
	// }

	result, err := elc.EvmLaunchpadService.GetMyAllLaunchpad(req.EoaAddress)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type tetetet struct {
	Name string `json:"name" binding:"required"`
}

func (elc *EvmLaunchpadController) RegisterEvmLaunchpadRoutes(r *gin.Engine) {
	route := r.Group("/evmLaunchpad")

	route.POST("/CreateNewLaunchPad", elc.CreateNewLaunchPad)
	route.GET("/GetMyAllLaunchpad", elc.GetMyAllLaunchpad)
}
