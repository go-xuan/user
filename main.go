package main

import (
	"github.com/go-xuan/quanx/app"
	"user/internal/model/entity"

	"user/internal/router"
)

func main() {
	var engine = app.NewEngine(
		app.EnableNacos,
		app.MultiRedis,
		app.MultiCache,
	)
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
