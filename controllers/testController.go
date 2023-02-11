package controllers

import (
	"net/http"

	"go_blockChain_server/models"
	"go_blockChain_server/services"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	TestService services.TestService
}

func New(t services.TestService) TestController {
	return TestController{
		TestService: t,
	}
}

func (tc *TestController) CreateTest(ctx *gin.Context) {
	var test models.Test

	if err := ctx.ShouldBindJSON(&test); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := tc.TestService.CreateTest(&test)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message ": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, test)
}

func (tc *TestController) GetAllTests(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "nil")
	// ctx.JSON(http.StatusOK, tc.TestService.GetTests())
}

func (tc *TestController) RegisterTestRoutes(r *gin.Engine) {
	testRoute := r.Group("/test")
	testRoute.POST("/create", tc.CreateTest)
	testRoute.GET("/getTests", tc.GetAllTests)
}
