package viper

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"offerBook/config"
)

func Init() error {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&config.Cfg); err != nil {
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		err := v.ReadInConfig()
		if err != nil {
			log.Fatal(err)
		}
	})
	return nil
}
