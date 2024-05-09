package models

import (
	"gorm.io/gorm"
)

type ToDoList struct {
	gorm.Model
	Owner              string            `json:"owner"`
	ListCompletionRate float64           `json:"listcompletionrate"`
	Messages           []ToDoListMessage `json:"messages"`
}

type ToDoListMessage struct {
	gorm.Model
	ToDoListID int    `json:"todolistid"`
	Content    string `json:"content"`
	IsItDone   bool   `json:"isitdone"`
}

// response for encapsulation of the TO-DO list and messages

type ToDoListResponse struct {
	Messages           []ToDoListMessage `json:"messages"`
	ListCompletionRate float64           `json:"listcompletionrate"`
}

type MessageResponse struct {
	ToDoListID int    `json:"todolistid"`
	Content    string `json:"content"`
	IsItDone   bool   `json:"isitdone"`
}
