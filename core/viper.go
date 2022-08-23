package core

import (
	"ToriBackend/env"
	"ToriBackend/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper() *viper.Viper {
	// parse the config file path from command line
	var configFile string
	if env.Mode == "debug" {
		configFile = "config_debug.yaml"
	} else if env.Mode == "release" {
		configFile = "config_release.yaml"
	} else {
		configFile = "config_debug.yaml"
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
