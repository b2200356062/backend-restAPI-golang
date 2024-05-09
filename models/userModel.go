package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Type     string // two types of users, first and second. first type is the default user, second type can see other users to-do lists.
	HasList  int
}
