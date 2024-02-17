package model

import (
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	gorm.Model
	Owner    int
	Status   int
	Company  string
	Position string
	Notes    string
	Public   bool
	Time     time.Time
}
