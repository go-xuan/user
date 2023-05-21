package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quanxiaoxuan/go-builder/database"
	"github.com/quanxiaoxuan/go-builder/logs"
	"github.com/quanxiaoxuan/go-builder/nacos"
	"github.com/quanxiaoxuan/go-builder/redis"
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
	logs.InitLogger(&conf.Config.Log, log.StandardLogger())

	// 初始化Nacos连接配置
	nacos.InitNacosConn(&conf.Config.Nacos)
	// 加载Nacos配置
	if nacos.InitConfigClient() {
		conf.LoadNacosConfig()
	}
	// 注册Nacos服务
	if nacos.InitNamingClient() {
		conf.InitNacosServerRegister()
	}

	// 连接数据库
	database.InitGormDB(&conf.Config.Database)
	// 初始化数据库表结构
	if conf.Config.Database.InitTable {
		common.InitGormTable()
	}
	// 连接redis
	redis.InitRedisConsole(&conf.Config.Redis)
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
	ginEngine.Use(logs.LoggerToFile(), gin.Recovery())
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
