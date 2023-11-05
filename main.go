package main

import (
	"github.com/go-xuan/quanx/engine"

	"quan-user/common"
	"quan-user/internal/router"
	"quan-user/model/table"
)

func main() {
	var newEngine = engine.GetEngine()
	newEngine.SetConfigPath("config/config-localhost.yaml")
	newEngine.AddModel(
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
