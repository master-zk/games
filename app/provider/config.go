package provider

import (
	"fmt"
	"games/app/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type configProvider struct {
}

var ConfigProvider = &configProvider{}

func (p *configProvider) Register() {
	ps := os.PathSeparator
	configPath := global.BasePath + string(ps) + "conf"
	fileName := getConfigFIleName()

	// 读取配置文件
	viper.AddConfigPath(configPath)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s \n", err)
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 重新映射配置到结构体
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&global.Config); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}

		// 重新设置数据库连接
		DatabaseProvider.Register()
	})

	// 将配置映射到结构体
	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func getConfigFIleName() string {
	configName := "config.yaml"

	switch os.Getenv("GAME_ENV") {
	case "production":
		configName = "config.prod.yaml"
	case "development":
		configName = "config.development.yaml"
	default:
		break
	}

	return configName
}
