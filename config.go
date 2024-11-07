package arpc

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		log.Println("Load config file from: ", configPath)
		viper.SetConfigFile(configPath)
	} else {
		log.Println("Load config file from: ./config.yaml")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

	// 设置默认值
	viper.SetDefault("grpc.addr", "0.0.0.0:8000")
	viper.SetDefault("http.addr", "0.0.0.0:8080")
	viper.SetDefault("grpc.max_msg_size", 100)

	// 加载配置文件
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
