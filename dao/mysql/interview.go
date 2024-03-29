package mysql

import (
	"goMian/model"
)

func (msq *mysqlDB) CreateInterview(it *model.Interview) error {
	err := msq.db.Create(&it).Error
	return err
}

func (msq *mysqlDB) FindLastInterview() (*model.Interview, error) {
	var it model.Interview
	err := msq.db.Model(&model.Interview{}).Order("id DESC").First(&it).Error
	return &it, err
}

func (msq *mysqlDB) FindLastInterviewDetail() (*model.InterviewDetail, error) {
	var detail model.InterviewDetail
	err := msq.db.Model(&model.InterviewDetail{}).Order("id DESC").First(&detail).Error
	return &detail, err
}

func (msq *mysqlDB) FindInterviewsByOwner(owner any) ([]model.Interview, error) {
	var its []model.Interview
	err := msq.db.Model(&model.Interview{}).Where("owner = ?", owner).Find(&its).Error
	return its, err
}

func (msq *mysqlDB) FindInterviewByID(id any) (*model.Interview, error) {
	var it model.Interview
	err := msq.db.Model(&model.Interview{}).Where("id = ?", id).First(&it).Error
	return &it, err
}

func (msq *mysqlDB) UpdateInterviewStatus(it *model.Interview) error {
	err := msq.db.Model(&model.Interview{}).Where("id = ?", it.ID).Update("status", it.Status).Error
	return err
}

func (msq *mysqlDB) UpdateInterviewDetail(it *model.Interview) error {
	err := msq.db.Model(&model.Interview{}).Where("id = ?", it.ID).Update("detail", it.Detail).Error
	return err
}

func (msq *mysqlDB) DeleteInterviewByID(id, itID any) error {
	err := msq.db.Where("owner = ?", id).Delete(&model.Interview{}, itID).Error
	return err
}

func (msq *mysqlDB) AddInterviewDetail(detail *model.InterviewDetail) error {
	err := msq.db.Create(detail).Error
	return err
}

func (msq *mysqlDB) DeleteDetailByID(id any) error {
	err := msq.db.Delete(&model.InterviewDetail{}, id).Error
	return err
}

func (msq *mysqlDB) InterviewStatus(it *model.Interview) error {
	if it.Time == 0 {
		return nil
	}
	it.InitStatus()
	err := msq.UpdateInterviewStatus(it)
	return err
}
