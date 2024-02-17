package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"index" binding:"min=4,max=16"`
	Password string `binding:"min=8,max=32"`
	Status   int
	Group    int
}
