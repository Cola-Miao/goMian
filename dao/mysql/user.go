package mysql

import (
	"errors"
	"goMian/model"
	"gorm.io/gorm"
)

func (msq *mysqlDB) CreateUser(u *model.User) error {
	err := msq.db.Create(&u).Error
	return err
}

func (msq *mysqlDB) UserExist(u *model.User) bool {
	if err := msq.db.First(&model.User{}, "user_name = ?", u.UserName).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (msq *mysqlDB) FindUserByName(name string) (*model.User, error) {
	var u *model.User
	err := msq.db.First(&u, "user_name = ?", name).Error
	return u, err
}
