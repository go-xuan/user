package main

import (
	"github.com/go-xuan/quanx"
	"user/internal/model/entity"

	"user/internal/router"
)

func main() {
	var engine = quanx.NewEngine(
		//quanx.EnableNacos(),   // 启用nacos
		quanx.MultiDatabase(), // 多数据源
		quanx.MultiRedis(),    // 多redis
		quanx.MultiCache(),    // 多缓存
		quanx.EnableQueue(),   // 使用队列
		quanx.EnableDebug(),   // 启用debug模式
		quanx.SetTable(
			&entity.User{},
			&entity.Role{},
			&entity.RoleUser{},
			&entity.Group{},
			&entity.GroupUser{},
			&entity.GroupRole{},
			&entity.Log{}),
		quanx.SetGinRouter(router.BindGinRouter),
	)
	// 初始化路由
	engine.RUN()
}
