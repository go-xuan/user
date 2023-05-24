package conf

import (
	"github.com/jinzhu/configor"
	"github.com/quanxiaoxuan/go-builder/gormx"
	"github.com/quanxiaoxuan/go-builder/logx"
	"github.com/quanxiaoxuan/go-builder/nacosx"
	"github.com/quanxiaoxuan/go-builder/redisx"
	"github.com/quanxiaoxuan/go-builder/snowflakex"
	"github.com/quanxiaoxuan/go-utils/ipx"

	"strconv"
)

var Config AppConfig

// 应用服务配置
type AppConfig struct {
	System   System                `json:"system" yaml:"system"`     // 应用配置
	Log      logx.Config           `json:"log" yaml:"log"`           // 日志配置
	Nacos    nacosx.Config         `json:"nacos" yaml:"nacos"`       // nacos访问配置
	Configs  nacosx.ConfigItemList `json:"configs" yaml:"configs"`   // nacos配置清单
	Database gormx.Config          `json:"database" yaml:"database"` // 数据库配置
	Redis    redisx.Config         `json:"redis" yaml:"redis"`       // redis配置
}

// 系统基础配置
type System struct {
	AppName string `json:"appName" yaml:"appName"`               // 应用名
	Host    string `json:"host" yaml:"host" default:"localhost"` // 主机IP
	Port    string `json:"port" yaml:"port" default:"8888"`      // 端口
	Env     string `json:"env" yaml:"env" default:"localhost"`   // 发布环境
	Prefix  string `json:"prefix" yaml:"prefix" default:"app"`   // 路由前缀
	Version string `json:"version" yaml:"version" default:"v1"`  // 版本
}

// 初始化本地应用配置
func InitAppConfig() {
	if err := configor.New(&configor.Config{
		Debug:       true,
		Environment: configor.ENV(),
	}).Load(&Config, "config.yml"); err != nil {
		panic(err)
	}
	if ipx.GetWLANIP() != "" {
		Config.System.Host = ipx.GetWLANIP()
	}
	Config.Log.LogName = Config.System.AppName
	// 初始化雪花ID生成器
	workerId, _ := strconv.ParseInt(Config.System.Host, 10, 64)
	snowflakex.InitSnowFlake(workerId % 1023)
}
