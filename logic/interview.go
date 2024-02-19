package logic

import (
	"goMian/dao/mysql"
	"goMian/dao/redis"
	"goMian/generator"
	"goMian/model"
)

func CreateInterview(it *model.Interview) error {
	id := <-generator.Gtr.InterviewID
	it.ID = uint(id)
	if err := mysql.DB.CreateInterview(it); err != nil {
		return err
	}
	if err := redis.DB.RelationInterview(it); err != nil {
		return err
	}
	return nil
}

func RefreshInterview(owner int) error {
	its, err := mysql.DB.FindInterviewsByOwner(owner)
	if err != nil {
		return err
	}
	var itsID []any
	for _, it := range its {
		itsID = append(itsID, int(it.ID))
	}
	if err = redis.DB.DeleteInterviewByOwner(owner); err != nil {
		return err
	}
	if err = redis.DB.FillInterview(itsID, owner); err != nil {
		return err
	}
	return nil
}
