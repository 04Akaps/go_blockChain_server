package controllers

import (
	"errors"
	"net/http"

	"go_blockChain_server/customerror"
	"go_blockChain_server/middleware"
	"go_blockChain_server/models"
	"go_blockChain_server/services"

	"go_blockChain_server/config"

	"github.com/gin-gonic/gin"
)

type EvmLaunchpadController struct {
	EvmLaunchpadService services.EvmLaunchpadService
	config              config.Config
}

func NewLaunchpadController(els services.EvmLaunchpadService, config config.Config) EvmLaunchpadController {
	return EvmLaunchpadController{
		EvmLaunchpadService: els,
		config:              config,
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

type getAllLaunchPadRequest struct {
	EoaAddress string `form:"eoa_address" binding:"startswith=0x"`
}

// @Summary getMy All EvmLaunchpad
// @Schemes
// @Description getMy All EvmLaunchpad
// @Tags EVM Launchpad
// @Produce json
// @Param eoa_address query string true "input my Eoa Address"
// @response default {object} []models.EvmLaunchpad{}
// @Router /evmLaunchpad/GetMyAllLaunchpad [get]
func (elc *EvmLaunchpadController) GetMyAllLaunchpad(ctx *gin.Context) {
	var req getAllLaunchPadRequest

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

type getLaunchPadRequest struct {
	CaAddress string `form:"ca_address" binding:"startswith=0x"`
}

// @Summary get EvmLaunchpad
// @Schemes
// @Description get EvmLaunchpad
// @Tags EVM Launchpad
// @Produce json
// @Param ca_address query string true "Input ca Address"
// @response default {object} models.EvmLaunchpad{}
// @Router /evmLaunchpad/GetLaunchPad [get]
func (elc *EvmLaunchpadController) GetOneLaunchpad(ctx *gin.Context) {
	var req getLaunchPadRequest

	queryCheckError := middleware.CheckUriQueryBinding(&req, ctx)
	if queryCheckError != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": queryCheckError, "status": -1})
		return
	}

	launchpad, err := elc.EvmLaunchpadService.GetOneLaunchpad(req.CaAddress)
	if err != nil {
		ctx.JSON(500, customerror.ErrorMsg(err))
		return
	}

	ctx.JSON(http.StatusOK, launchpad)
}

func (elc *EvmLaunchpadController) DeleteAllLaunchpadByAdmin(ctx *gin.Context) {
	// debug용 후에 데이터 전체 삭제해야 할 떄 사용
	password := ctx.Request.Header["Password"]

	if len(password) == 0 {
		ctx.AbortWithError(http.StatusMethodNotAllowed, errors.New("password is not exist"))
		return
	}

	if elc.config.AdminPassword == ctx.Request.Header["Password"][0] {
		elc.EvmLaunchpadService.DeleteAllLaunchpadByAdmin()

		ctx.JSON(http.StatusOK, gin.H{"status": 0, "message": "success"})
	} else {
		ctx.AbortWithError(http.StatusMethodNotAllowed, errors.New("password error	"))
		return
	}
}

func (elc *EvmLaunchpadController) RegisterEvmLaunchpadRoutes(r *gin.Engine) {
	route := r.Group("/evmLaunchpad")

	route.POST("/CreateNewLaunchPad", elc.CreateNewLaunchPad)
	route.GET("/GetMyAllLaunchPad", elc.GetMyAllLaunchpad)
	route.GET("/GetLaunchPad", elc.GetOneLaunchpad)
	route.DELETE("/DeleteAllLaunchpadData", elc.DeleteAllLaunchpadByAdmin)
}
