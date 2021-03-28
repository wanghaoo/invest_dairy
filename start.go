package main

import (
	"fmt"
	"invest_dairy/common"
	_ "invest_dairy/controller"
	_ "invest_dairy/controller/admin"
	"invest_dairy/service"
	"invest_dairy/util"

	"github.com/yb7/echoswg"
)

func main1() {
	common.OpenDB()
	common.OpenRedis()
	defer common.CloseRedis()
	service.InitOssBucket()

	e := util.EchoInst
	e.Static("/swagger", "swagger")
	e.Static("/contract", "contract")
	e.GET("/api-docs", echoswg.GenApiDoc("Server", "Server-API"))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", "1111")))
}
