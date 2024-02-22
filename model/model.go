package model

import (
	"goMian/config/inner"
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	gorm.Model `json:"-"`
	Owner      int    `json:"-"`
	Status     int    `json:"-"` // -2 Not yet, -1 Close, 0, 1 Over
	Detail     int    `json:"-"`
	Time       int64  `json:"time"`
	Company    string `json:"company,omitempty"`
	Position   string `json:"position,omitempty"`
	Memo       string `json:"memo,omitempty"`
	Public     bool   `json:"public,omitempty"`
}

type InterviewDetail struct {
	gorm.Model     `json:"-"`
	Question       string `json:"question"`
	Answer         string `json:"answer"`
	Notes          string `json:"notes,omitempty"`
	SelfEvaluation string `json:"self-evaluation,omitempty"`
	HREvaluation   string `json:"hr-evaluation,omitempty"`
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
