package main

import (
	"context"
	"database/sql"
	"log"

	"go_blockChain_server/controllers"
	"go_blockChain_server/models"
	migrate "go_blockChain_server/mysql"
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

	// evmLaunchpad mysql

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/launchPad?multiStatements=true&parseTime=true")
	if err != nil {
		log.Fatal("launchpad sql Open Error : ", err)
	}

	launchpadCtx := context.Background()
	query := migrate.MigratMysql(db)

	els := services.NewEvmLaunchpadServiceImpl(launchpadCtx, query)
	elc := controllers.NewLaunchpadController(els)

	elc.RegisterEvmLaunchpadRoutes(server)
	tc.RegisterTestRoutes(server)

	err = server.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
