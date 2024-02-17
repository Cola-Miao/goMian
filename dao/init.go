package dao

import (
	"offerBook/dao/mysql"
)

func Init() (err error) {
	if err = mysql.DB.Init(); err != nil {
		return err
	}
	//if err = redis.DB.Init(); err != nil {
	//	return err
	//}
	return err
}
