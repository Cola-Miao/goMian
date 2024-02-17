package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"offerBook/pkg/log"
	"offerBook/pkg/viper"
)

func Init() (err error) {
	if err = log.Init(); err != nil {
		return
	}
	if err = viper.Init(); err != nil {
		return
	}
	jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name})
	return
}
