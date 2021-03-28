package main

import (
	"fmt"
	"invest_dairy/common"
	_ "invest_dairy/controller"
	_ "invest_dairy/controller/admin"
	"invest_dairy/service"
	"invest_dairy/util"

	"github.com/labstack/echo/v4/middleware"
	"github.com/yb7/echoswg"
)

func main() {
	common.OpenDB()
	common.OpenRedis()
	defer common.CloseRedis()
	service.InitOssBucket()
	go service.InitCron()

	e := util.EchoInst
	e.Static("/swagger", "swagger")
	e.Static("/contract", "contract")
	e.GET("/api-docs", echoswg.GenApiDoc("Server", "Server-API"))
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", "80")))
}
