package controllers

import (
	"net/http"

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
	ctx.JSON(http.StatusOK, nil)
}

func (elc *EvmLaunchpadController) GetMyAllLaunchpad(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (elc *EvmLaunchpadController) GetMyLaunchpad(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (elc *EvmLaunchpadController) RegisterEvmLaunchpadRoutes(r *gin.Engine) {
	route := r.Group("/evmLaunchpad")

	route.POST("/CreateNewLaunchPad", elc.CreateNewLaunchPad)
	route.POST("/GetMyAllLaunchpad", elc.GetMyAllLaunchpad)
	route.POST("/GetMyLaunchpad", elc.GetMyLaunchpad)
}
