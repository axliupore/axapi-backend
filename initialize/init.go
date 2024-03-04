package initialize

import (
	"fmt"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化所有的配置信息
func InitConfig() {
	// 读取配置文件
	readConfig()
	// 初始化 zap
	InitZap()
	// 初始化 mysql
	InitMysql()
	// 初始化 redis
	InitRedis()
	if global.Config.Server.UseOSS {
		// 初始化 aliyun-oss
		InitAliyunOSS()
	}
	// 初始化路由
	InitRouter()
}

// 读取配置文件
func readConfig() {
	v := viper.New()
	// todo 这里记得改为 config.yaml
	v.SetConfigFile("config.private.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig() // 在配置文件变更时自动加载配置
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	if global.Config.Server.Mode == "local" {
		fmt.Println("read config file:")
		spew.Dump(global.Config)
	}
}
