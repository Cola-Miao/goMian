package logic

import (
	"goMian/config/inner"
	"goMian/dao/mysql"
	"goMian/dao/redis"
	"goMian/generator"
	"goMian/model"
	"golang.org/x/sync/errgroup"
)

func CreateInterview(it *model.Interview) error {
	id := <-generator.Gtr.InterviewID
	it.ID = uint(id)
	it.InitStatus()
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
	if len(its) < 1 {
		return nil
	}
	var eg errgroup.Group
	eg.SetLimit(inner.ErrorGroupLimit)
	for _, it := range its {
		cp := it
		eg.TryGo(func() error {
			err = mysql.DB.InterviewStatus(&cp)
			return err
		})
		itsID = append(itsID, int(it.ID))
	}
	if err = eg.Wait(); err != nil {
		return err
	}
	if err = redis.DB.DeleteInterviewByOwner(owner); err != nil {
		return err
	}
	if err = redis.DB.FillInterview(itsID, owner); err != nil {
		return err
	}
	return nil
}
