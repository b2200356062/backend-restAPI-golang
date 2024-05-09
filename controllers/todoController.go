package controllers

import (
	"fmt"
	"net/http"
	"staj/initializers"
	"staj/models"
	"time"

	"github.com/gin-gonic/gin"
)

//TODO: JWT İLE AUTHORİZATİON SAĞLA

func CreateTODOList(c *gin.Context) {

	var body struct {
		Owner              string
		ListCompletionRate float64
		Messages           []models.ToDoListMessage
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read TO-DO body",
		})
		return
	}

	todolist := models.ToDoList{Owner: body.Owner, ListCompletionRate: body.ListCompletionRate, Messages: body.Messages}
	result := initializers.DB.Create(&todolist)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create TO-DO List",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "TO-DO list created successfully",
	})

}

func GetToDoLists(c *gin.Context) {

	var todolists []models.ToDoList

	// returns todo lists with messages loaded

	result := initializers.DB.Model(&models.ToDoList{}).Preload("Messages").Find(&todolists).Error

	if result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve TO-DO lists",
		})
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"todolists": todolists,
	})
}

func DeleteToDoList(c *gin.Context) {

	var todolists models.ToDoList
	var messages []models.ToDoListMessage

	todolistID := c.Param("id")

	//FIXME: update ve delete zamanı milimlerle farklı

	// update modify time, update delete time, but not actually delete the record
	initializers.DB.Model(&todolists).Where("id = ?", todolistID).Update("deleted_at", time.Now())

	// same for messages for the todolist
	initializers.DB.Model(&messages).Where("to_do_list_id = ?", todolistID).Update("deleted_at", time.Now())

	// success message
	c.JSON(200, gin.H{
		"message": "ToDo List deleted successfully",
	})
}

func CreateToDoMessage(c *gin.Context) {

	var body struct {
		ToDoListID int
		Content    string
		IsItDone   bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read TO-DO message body",
		})
		fmt.Print(c.Errors)
		return
	}

	message := models.ToDoListMessage{ToDoListID: body.ToDoListID, Content: body.Content, IsItDone: body.IsItDone}
	result := initializers.DB.Create(&message)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create TO-DO List",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("TO-DO message to list {%d} is created successfully", message.ToDoListID),
	})

}

func DeleteToDoMessage(c *gin.Context) {

	var message models.ToDoListMessage

	messageID := c.Param("id")

	// update modify time, update delete time, but not actually delete the record
	initializers.DB.Model(&message).Where("id = ?", messageID).Update("deleted_at", time.Now())

	// success message
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Message with ID {%s} deleted", messageID),
	})
}

func UpdateToDoMessage() {

}
