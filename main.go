package main

import (
	"github.com/go-xuan/quanx/app"

	"user/internal/model/entity"
	"user/internal/router"
)

func main() {
	var engine = app.NewEngine(
		app.EnableNacos,   // 启用nacos
		app.MultiDatabase, // 多数据源
		app.MultiRedis,    // 对redis
		app.MultiCache,    // 多缓存
		app.UseQueue,      // 使用队列
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
