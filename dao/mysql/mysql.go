package mysql

import (
	"fmt"
	"goMian/config"
	"goMian/config/inner"
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
	if err := createDatabase(config.Cfg.Mysql.Database); err != nil {
		return err
	}
	db, err := openDatabase()
	if err != nil {
		return err
	}
	if err = db.AutoMigrate(migrateList...); err != nil {
		return err
	}
	msq.db = db
	return nil
}

func createDatabase(dbname string) error {
	dsn := fmt.Sprintf(inner.DsnLite,
		config.Cfg.Mysql.Name,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Addr,
		config.Cfg.Mysql.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname).Error; err != nil {
		return err
	}
	return nil
}

func openDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(inner.Dsn,
		config.Cfg.Mysql.Name,
		config.Cfg.Mysql.Password,
		config.Cfg.Mysql.Addr,
		config.Cfg.Mysql.Port,
		config.Cfg.Mysql.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}
