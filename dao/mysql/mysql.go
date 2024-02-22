package mysql

import (
	"fmt"
	"goMian/config"
	"goMian/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = new(mysqlDB)

type mysqlDB struct {
	db *gorm.DB
}

var migrateList = []any{
	&model.User{},
	&model.Interview{},
	&model.InterviewDetail{},
}

func (msq *mysqlDB) Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Cfg.Mysql.Name,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Addr,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	if err = db.AutoMigrate(migrateList...); err != nil {
		return err
	}
	msq.db = db
	return nil
}
