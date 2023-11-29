package main

import (
	"github.com/go-xuan/quanx"

	"user/internal/model/table"
	"user/internal/router"
)

func main() {
	var engine = quanx.GetEngine()
	//engine.SetEngineConfig("config/config-localhost.yaml")
	engine.SetEngineConfig("conf/config.yaml")
	engine.AddModel(
		&table.User{},
		&table.UserAuth{},
		&table.Role{},
		&table.RoleUser{},
		&table.Group{},
		&table.GroupUser{},
		&table.GroupRole{},
		&table.Log{},
	)
	// 初始化路由
	engine.AddGinRouter(router.BindGinRouter)
	engine.RUN()
}
