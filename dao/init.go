package dao

import (
	"goMian/dao/mysql"
	"goMian/dao/redis"
)

func Init() (err error) {
	if err = mysql.DB.Init(); err != nil {
		return err
	}
	if err = redis.DB.Init(); err != nil {
		return err
	}
	return err
}
