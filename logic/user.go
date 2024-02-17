package logic

import (
	"errors"
	"offerBook/dao/mysql"
	"offerBook/model"
	"offerBook/utils"
)

func NewUser(u *model.User) error {
	if exist := mysql.DB.UserExist(u); exist {
		return errors.New("username exist")
	}
	pwd, err := utils.Encrypt([]byte(u.Password))
	if err != nil {
		return err
	}
	u.Password = string(pwd)
	err = mysql.DB.CreateUser(u)
	return err
}

func Login(u *model.User) (token string, err error) {
	us, err := mysql.DB.FindUserByName(u.UserName)
	if err != nil {
		return
	}
	if err = utils.CompareEncode([]byte(us.Password), []byte(u.Password)); err != nil {
		return
	}
	token, err = utils.GenerateJWT(int(us.ID))
	return
}
