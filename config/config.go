package config

import (
	"fmt"
	"games/app/global"
	"games/app/initialize"
	"github.com/fsnotify/fsnotify"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/spf13/viper"
	"log"
	"os"
)

func LoadGlobalConfig() {

	global.BasePath = gfile.Pwd()

	ps := os.PathSeparator
	configPath := global.BasePath + string(ps) + "config"
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

		initialize.OnConfigChangeInit()
	})

	// 将配置映射到结构体
	if err := viper.Unmarshal(&global.Config); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func getConfigFIleName() string {
	return "config.yaml"
}
