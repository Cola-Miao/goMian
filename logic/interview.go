package logic

import (
	"errors"
	rdb "github.com/redis/go-redis/v9"
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

func RefreshInterview(owner any) error {
	its, err := mysql.DB.FindInterviewsByOwner(owner)
	if err != nil {
		return err
	}
	var itsZ []rdb.Z
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
		itsZ = append(itsZ, rdb.Z{
			Score:  float64(it.Time),
			Member: it.ID,
		})
	}
	if err = eg.Wait(); err != nil {
		return err
	}
	if err = redis.DB.DeleteInterviewByOwner(owner); err != nil {
		return err
	}
	if err = redis.DB.FillInterview(itsZ, owner); err != nil {
		return err
	}
	return nil
}

func InterviewList(id int) ([]model.Interview, error) {
	itsID, err := redis.DB.FindInterviewsByOwner(id)
	if err != nil {
		return nil, err
	}
	var its []model.Interview
	for _, itID := range itsID {
		var it *model.Interview
		it, err = mysql.DB.FindInterviewByID(itID)
		if err != nil {
			return nil, err
		}
		its = append(its, *it)
	}
	return its, nil
}

func DeleteInterview(id int, itID string) error {
	if err := mysql.DB.DeleteInterviewByID(id, itID); err != nil {
		return err
	}
	if err := redis.DB.DeleteInterviewByID(id, itID); err != nil {
		return err
	}
	return nil
}

func AddInterviewDetail(id int, relevance string, detail *model.InterviewDetail) error {
	it, err := mysql.DB.FindInterviewByID(relevance)
	if err != nil {
		return err
	}
	if id != it.Owner {
		return errors.New("no auth")
	}
	if it.Detail > 0 {
		if err = mysql.DB.DeleteDetailByID(it.Detail); err != nil {
			return err
		}
	}
	detailID := <-generator.Gtr.DetailID
	detail.ID = uint(detailID)
	it.Detail = detailID
	if err = mysql.DB.UpdateInterviewDetail(it); err != nil {
		return err
	}
	if err = mysql.DB.AddInterviewDetail(detail); err != nil {
		return err
	}
	return nil
}
