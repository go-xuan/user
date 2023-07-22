package main

import (
	"github.com/quanxiaoxuan/quanx/server"
	"quan-admin/app/handler"
)

func main() {
	var engine = server.NewEngine()
	engine.AddMiddleware(
		engine.InitLogger,
		engine.InitNacos,
		engine.InitGorm)
	engine.AddRouterLoaders(handler.LoadApiRouter)
	engine.Run()
}
