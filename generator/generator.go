package generator

import (
	"errors"
	"goMian/dao/mysql"
	"gorm.io/gorm"
)

var Gtr Generator

type Generator struct {
	InterviewID chan int
}

func (g *Generator) Init() error {
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
