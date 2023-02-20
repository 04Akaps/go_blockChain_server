package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"go_blockChain_server/config"
	"go_blockChain_server/controllers"
	"go_blockChain_server/models"
	"go_blockChain_server/services"

	migrate "go_blockChain_server/mysql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var err error

var (
	testCtx  context.Context
	ts       services.TestService
	tc       controllers.TestController
	testList []*models.Test
)

var configType config.Config

func init() {
	// Test Set
	testCtx = context.TODO()
	testList = []*models.Test{{"test1", 1}, {"test2", 3}}

	ts = services.NewTestService(testList, testCtx)
	tc = controllers.New(ts)

	gin.DisableConsoleColor()

	// --- Create Log Files ---
	// t := time.Now()
	// startTime := t.Format("2006-01-02 15:04:05")
	// logFile := "logList/server_log -" + startTime + ".log"

	// _, err := os.Create(logFile)
	if err != nil {
		log.Fatal(err)
	}

	configType = config.LoadConfig(".")

	// gin.DefaultWriter = io.MultiWriter(f)
}

const (
	hashSample      = "0x000f62fb7dfc9ee08f93b8cf1f8c255b39da541e8d640f9cfd85d1d1857ca8a7"
	signedPublicKey = "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5"
)

func main() {
	// gin.SetMode(gin.ReleaseMode) // 후에 필요할 떄 사용
	server := gin.Default()

	db, err := sql.Open("mysql", configType.DbUri)
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

	client, err := ethclient.DialContext(launchpadCtx, configType.InfuraURL)
	if err != nil {
		log.Fatal("ethClient DialContext failed", err)
	}

	txHash := common.HexToHash(hashSample)

	tx, _, _ := client.TransactionByHash(context.Background(), txHash)
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal("=------%v\n", err)
	}
	result := msg.From()
	fmt.Printf("%v\n", result)

	// v, _, _ := tx.RawSignatureValues()
	// // receipt, err := client.TransactionReceipt(launchpadCtx, txHash)
	// // if err != nil {
	// // 	log.Fatal("=------%v\n", err)
	// // }

	// fmt.Printf(v)

	// // fmt.Println(receipt)
	// result, err := crypto.Ecrecover(txHash.Bytes(), append(tx.Data(), v.Bytes()...))
	// if err != nil {
	// 	log.Fatal("=-asdsd----%v\n", err)
	// }
	// fmt.Println(result)

	els := services.NewEvmLaunchpadServiceImpl(launchpadCtx, query)
	elc := controllers.NewLaunchpadController(els, configType, client)

	controllers.SwaggerSet(server)
	elc.RegisterEvmLaunchpadRoutes(server)
	tc.RegisterTestRoutes(server)

	err = server.Run(configType.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	defer client.Close()
}
