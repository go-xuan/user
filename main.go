package main

import (
	"github.com/quanxiaoxuan/quanx/engine"

	"quan-admin/common"
	"quan-admin/internal/router"
	"quan-admin/model/table"
)

func main() {
	var newEngine = engine.GetEngine()
	//newEngine.SetConfigPath("config/config.yaml")
	newEngine.SetConfigPath("config/config-localhost.yaml")
	newEngine.AddGormModel("default",
		&table.User{},
		&table.UserAuth{},
		&table.Role{},
		&table.RoleUser{},
		&table.Group{},
		&table.GroupUser{},
		&table.GroupRole{},
		&table.Log{},
	)
	// 初始化配置
	newEngine.AddInitializer(common.Init)
	// 初始化路由
	newEngine.AddRouterLoaders(router.BindGinRouter)
	newEngine.RUN()
}
