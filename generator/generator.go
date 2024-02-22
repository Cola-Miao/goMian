package generator

import (
	"errors"
	"goMian/dao/mysql"
	"gorm.io/gorm"
)

var Gtr Generator

type Generator struct {
	InterviewID chan int
	DetailID    chan int
}

func (g *Generator) Init() error {
	if err := g.initInterviewGenerator(); err != nil {
		return err
	}
	if err := g.initDetailGenerator(); err != nil {
		return err
	}
	return nil
}

func (g *Generator) initDetailGenerator() error {
	g.DetailID = make(chan int)
	detail, err := mysql.DB.FindLastInterviewDetail()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	go func() {
		id := int(detail.ID)
		for {
			id++
			g.DetailID <- id
		}
	}()
	return nil
}

func (g *Generator) initInterviewGenerator() error {
	g.InterviewID = make(chan int)
	it, err := mysql.DB.FindLastInterview()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	go func() {
		id := int(it.ID)
		for {
			id++
			g.InterviewID <- id
		}
	}()
	return nil
}
