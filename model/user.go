package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	UserName   string `json:"username" gorm:"index" binding:"min=4,max=16"`
	Password   string `json:"password" binding:"min=8,max=32"`
	Status     int    `json:"-"`
	Group      int    `json:"-"`
}
