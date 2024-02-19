package model

import (
	"gorm.io/gorm"
	"time"
)

type Interview struct {
	gorm.Model
	Owner    int
	Status   int
	Company  string        `json:"company,omitempty"`
	Position string        `json:"position,omitempty"`
	Notes    string        `json:"notes,omitempty"`
	Public   bool          `json:"public,omitempty"`
	Time     time.Duration `json:"time"`
}
