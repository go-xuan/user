package main

import (
	"github.com/quanxiaoxuan/quanx/server"
	"quan-admin/app/handler"
	"quan-admin/common"
)

func main() {
	var engine = server.NewEngine()
	engine.AddMiddleware(
		engine.InitLogger,
		engine.InitNacos,
		engine.InitGorm,
		common.InitGormTable)
	engine.AddRouterLoaders(handler.LoadApiRouter)
	engine.Run()
}
