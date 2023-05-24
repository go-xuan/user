package conf

import (
	"github.com/quanxiaoxuan/go-builder/nacosx"
	log "github.com/sirupsen/logrus"
)

// 初始化Nacos服务注册配置
func InitNacosServerRegister() {
	nacosx.InitNacosServerInstance(&nacosx.ServerConfig{
		Group: Config.System.Env,
		Name:  Config.System.AppName,
		Ip:    Config.System.Host,
		Port:  Config.System.Port,
	})
}

// 加载Nacos配置
func LoadNacosConfig() {
	if Config.Configs != nil && len(Config.Configs) > 0 {
		var err error
		var system nacosx.ConfigItem             // 系统配置
		var customConfList nacosx.ConfigItemList // 自定义通用配置
		for _, config := range Config.Configs {
			if config.Name == Config.System.AppName {
				system = config
			} else {
				customConfList = append(customConfList, config)
			}
		}
		// 初始化nacos配置监控
		nacosx.InitNacosConfigMonitor()
		// 加载通用配置
		err = customConfList.LoadNacosConfigBatch(&Config)
		if err != nil {
			log.Error("加载Nacos通用配置失败！", err)
		}
		// 加载系统配置
		err = system.LoadNacosConfig(&Config)
		if err != nil {
			log.Error("加载Nacos系统配置失败!", err)
		}
	}
}
