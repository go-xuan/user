package main

import (
	"github.com/go-xuan/quanx/server"

	"user/internal/model/entity"
	"user/internal/router"
)

func main() {
	var engine = server.GetEngine(server.EnableNacos)
	engine.AddGinRouter(router.BindGinRouter)
	engine.AddTable(
		&entity.User{},
		&entity.UserAuth{},
		&entity.Role{},
		&entity.RoleUser{},
		&entity.Group{},
		&entity.GroupUser{},
		&entity.GroupRole{},
		&entity.Log{},
	)
	// 初始化路由
	engine.RUN()
}
