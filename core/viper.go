package core

import (
	"WebAppStructure/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func InitViper() *viper.Viper {
	// parse the config file path from command line
	var configFile string
	flag.StringVar(&configFile, "configFile", "", "the path of config yaml file")
	flag.Parse()
	if configFile == "" { // 优先级：命令行 > 环境变量 > 默认值
		if envConfigFile := os.Getenv("CONFIG_FILE"); envConfigFile == "" {
			configFile = "config_debug.yaml"
		} else {
			configFile = envConfigFile
		}
	}

	v := viper.New()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error: viper read config file failed: %s \n", err))
	}

	//useWatchConfig(v)
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	return v
}

func useWatchConfig(v *viper.Viper) {
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
		if err := v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
}
