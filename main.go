package main

import (
	"context"
	"log"

	"go_blockChain_server/controllers"
	"go_blockChain_server/models"
	"go_blockChain_server/services"

	"github.com/gin-gonic/gin"
)

var err error

var (
	testCtx  context.Context
	ts       services.TestService
	tc       controllers.TestController
	testList []*models.Test
)

func init() {
	// Test Set
	testCtx = context.TODO()
	testList = []*models.Test{{"test1", 1}, {"test2", 3}}

	ts = services.NewTestService(testList, testCtx)
	tc = controllers.New(ts)
}

func main() {
	server := gin.Default()

	tc.RegisterTestRoutes(server)

	err := server.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}
