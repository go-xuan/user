package main

import (
	"github.com/go-xuan/quanx"
	"user/internal/router"
)

func main() {
	var engine = quanx.GetEngine(quanx.EnableNacos)
	engine.AddGinRouter(router.BindGinRouter)
	//engine.SetConfigDir("conf")
	//engine.AddTable(
	//	&table.User{},
	//	&table.UserAuth{},
	//	&table.Role{},
	//	&table.RoleUser{},
	//	&table.Group{},
	//	&table.GroupUser{},
	//	&table.GroupRole{},
	//	&table.log{},
	//)
	// 初始化路由
	engine.RUN()
}
