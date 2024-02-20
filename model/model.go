package model

import (
	"goMian/config/inner"
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	gorm.Model
	Owner    int
	Status   int    // -2 Not yet, -1 Close, 0, 1 Over
	Time     int64  `json:"time"`
	Company  string `json:"company,omitempty"`
	Position string `json:"position,omitempty"`
	Notes    string `json:"notes,omitempty"`
	Public   bool   `json:"public,omitempty"`
}

func (it *Interview) InitStatus() {
	now := time.Now()
	itm := time.Unix(it.Time, 0)
	switch {
	case now.After(itm.Add(inner.InterviewExpiresTime)):
		it.Status = 1
	case now.After(itm), now.Equal(itm):
		it.Status = 0
	case now.Before(itm.Add(-inner.InterviewBufferTime)):
		it.Status = -1
	default:
		it.Status = -2
	}
}
