package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/gormx"
	"github.com/quanxiaoxuan/go-builder/logx"
	"github.com/quanxiaoxuan/go-builder/nacosx"
	"github.com/quanxiaoxuan/go-builder/redisx"
	log "github.com/sirupsen/logrus"

	"quan-admin/app/handler"
	"quan-admin/common"
	"quan-admin/conf"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("服务启动失败，错误为 : %s", err)
			return
		}
		select {}
	}()
	// 初始化本地配置文件
	conf.InitAppConfig()

	// 初始化日志
	logx.InitLogger(&conf.Config.Log)

	// 初始化Nacos连接配置
	nacosx.InitNacosConn(&conf.Config.Nacos)
	// 加载Nacos配置
	if nacosx.InitConfigClient() {
		conf.LoadNacosConfig()
	}
	// 注册Nacos服务
	if nacosx.InitNamingClient() {
		conf.InitNacosServerRegister()
	}

	// 初始化Gorm
	gormx.InitGormCTL(&conf.Config.Database)
	// 初始化数据库表结构
	if conf.Config.Database.InitTable {
		common.InitGormTable()
	}
	// 初始化redis
	redisx.InitRedisCTL(&conf.Config.Redis)
	// 初始化gin路由
	InitGinRouter(&conf.Config.System)

	log.Info("服务启动成功！！！")
}

// 初始化gin路由
func InitGinRouter(sys *conf.System) {
	if sys.Env == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	ginEngine := gin.New()
	ginEngine.Use(logx.LoggerToFile(), gin.Recovery())
	_ = ginEngine.SetTrustedProxies([]string{sys.Host})
	// 注册路由
	router := ginEngine.Group(sys.Prefix).Group(sys.Version)
	handler.AddHandlerFunc(router)
	log.Info("=== API接口请求地址: http://" + sys.Host + ":" + sys.Port)
	// 监听端口
	if err := ginEngine.Run(":" + sys.Port); err != nil {
		panic(err)
	}
}
