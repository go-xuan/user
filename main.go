package main

import (
	"github.com/go-xuan/quanx"

	"user/common"
	"user/internal/router"
)

func main() {
	var engine = quanx.GetEngine()
	//engine.SetConfigPath("config/config-localhost.yaml")
	engine.SetConfigPath("config/config.yaml")
	//engine.AddModel(
	//	&table.User{},
	//	&table.UserAuth{},
	//	&table.Role{},
	//	&table.RoleUser{},
	//	&table.Group{},
	//	&table.GroupUser{},
	//	&table.GroupRole{},
	//	&table.Log{},
	//)
	// 初始化配置
	engine.AddInitializer(common.Init)
	// 初始化路由
	engine.AddRouterLoaders(router.BindGinRouter)
	engine.RUN()
}
