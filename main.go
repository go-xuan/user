package main

import (
	"github.com/go-xuan/quanx"
	"user/internal/model/table"
	"user/internal/router"
)

func main() {
	var engine = quanx.GetEngine(quanx.UseNacos)
	engine.AddGinRouter(router.BindGinRouter)
	//engine.SetConfig("conf/config-localhost.yaml")
	engine.AddTable(
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
	engine.RUN()
}
