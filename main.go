package main

import (
	"github.com/go-xuan/quanx/core"
	"user/internal/model/entity"
	"user/internal/router"
)

func main() {
	var engine = core.GetEngine(core.EnableNacos)
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
