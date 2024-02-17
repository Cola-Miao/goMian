package mysql

import "offerBook/model"

func (msq *mysqlDB) CreateInterview(it *model.Interview) error {
	err := msq.db.Create(&it).Error
	return err
}
