package main

import (
	"context"
	"database/sql"
	"log"
	"time"

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
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// evmLaunchpad mysql

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/launchPad?multiStatements=true&parseTime=true")
	if err != nil {
		log.Fatal("launchpad sql Open Error : ", err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(6)
	// 인당 최대 connection수 제한
	// 한명이 여러개의 db에 접근할 이유는 없다고 생각하기 떄문에 3으로 설정
	// 이러한 설정을 통해서 connection이 재 사용 된다.
	// sqlc 를 사용하였기 떄문에 어차피 query날리고 해당 함수에서 close를 실행함 -> connection재활용 가능

	launchpadCtx := context.Background()
	query := migrate.MigratMysql(db)

	els := services.NewEvmLaunchpadServiceImpl(launchpadCtx, query)
	elc := controllers.NewLaunchpadController(els)

	elc.RegisterEvmLaunchpadRoutes(server)
	tc.RegisterTestRoutes(server)

	err = server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
