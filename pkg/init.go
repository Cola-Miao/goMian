package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"goMian/pkg/cron"
	"goMian/pkg/log"
	"goMian/pkg/viper"
)

func Init() (err error) {
	if err = log.Init(); err != nil {
		return
	}
	if err = viper.Init(); err != nil {
		return
	}
	if err = cron.Init(); err != nil {
		return err
	}
	jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})
	return
}
