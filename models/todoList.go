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

// response for encapsulation of the todolist

// type ToDoListResponse struct {
// 	// returns the fields below
// 	Owner              string            `json:"owner"`
// 	Messages           []ToDoListMessage `json:"messages"`
// 	ListCompletionRate float64           `json:"listcompletionrate"`
// 	CreatedAt          time.Time         `json:"created_at"`
// 	UpdatedAt          time.Time         `json:"updated_at"`
// 	DeletedAt          time.Time         `json:"deleted_at"`
// }

// // using the response struct
// var response []models.ToDoListResponse

// for _, todo := range todolists {

// 	// if the record is not deleted, it is visible. however, if it seems deleted,
// 	// its hidden in the response but visible in the database

// 	if !todo.DeletedAt.Time.IsZero() {
// 		continue
// 	}

// 	response = append(response, models.ToDoListResponse{
// 		ID:                 todo.ID,
// 		Owner:              todo.Owner,
// 		Messages:           todo.Messages,
// 		CreatedAt:          todo.CreatedAt,
// 		UpdatedAt:          todo.UpdatedAt,
// 		DeletedAt:          todo.DeletedAt.Time,
// 		ListCompletionRate: todo.ListCompletionRate,
// 	})
// }
