package main

import (
	"github.com/go-xuan/quanx"
	"user/internal/model/entity"
	"user/internal/router"
)

func main() {
	var engine = quanx.NewEngine(
		//app.EnableNacos,   // 启用nacos
		//app.MultiDatabase, // 多数据源
		quanx.MultiRedis,  // 对redis
		quanx.MultiCache,  // 多缓存
		quanx.EnableQueue, // 使用队列
		quanx.Debug,       // debug模式
	)

	engine.AddGinRouter(router.BindGinRouter)
	engine.AddTable(
		&entity.User{},
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
