package logic

import (
	"offerBook/dao/mysql"
	"offerBook/model"
)

func CreateInterview(it *model.Interview) error {
	err := mysql.DB.CreateInterview(it)
	return err
}
