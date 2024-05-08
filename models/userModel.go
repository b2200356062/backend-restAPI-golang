package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Type     string // two types of users, default and super. super users can see all users lists
}
