package models

import (
	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	Owner              string            `json:"list owner"`
	ListCompletionRate float64           `json:"list completion rate"`
	Messages           []ToDoListMessage `json:"messages"`
}

type ToDoListMessage struct {
	gorm.Model
	ToDoListID uint   `json:"to do list id"`
	Content    string `json:"content"`
	IsItDone   bool   `json:"is it done"`
}

// response for encapsulation of the TO-DO list and messages
// i purposefully hid the time stamps from users. it can be made visible with just a few lines.
type ToDoListResponse struct {
	Owner              string            `json:"list owner"`
	Messages           []MessageResponse `json:"messages"`
	ListCompletionRate float64           `json:"list completion rate"`
}

type MessageResponse struct {
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	IsItDone bool   `json:"is it done"`
}
