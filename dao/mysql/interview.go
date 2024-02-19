package mysql

import "goMian/model"

func (msq *mysqlDB) CreateInterview(it *model.Interview) error {
	err := msq.db.Create(&it).Error
	return err
}

func (msq *mysqlDB) FindLastInterview() (*model.Interview, error) {
	var it model.Interview
	err := msq.db.Model(&model.Interview{}).Order("id DESC").First(&it).Error
	return &it, err
}

func (msq *mysqlDB) FindInterviewsByOwner(owner int) ([]model.Interview, error) {
	var its []model.Interview
	err := msq.db.Model(&model.Interview{}).Where("owner = ?", owner).Find(&its).Error
	return its, err
}
