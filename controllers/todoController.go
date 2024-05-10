package controllers

import (
	"fmt"
	"net/http"
	"staj/initializers"
	"staj/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTODOList(c *gin.Context) {

	var body struct {
		ListCompletionRate float64
		Messages           []models.ToDoListMessage
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read TO-DO body",
		})
		return
	}

	// gets current user name from validation
	owner := getCurrentUser(c)

	// each user can have 1 TO-DO list
	if owner.HasList != 0 {
		c.JSON(http.StatusNotAcceptable, "Each user can only have 1 TO-DO list")
		return
	}

	todolist := models.ToDoList{Owner: owner.Name, ListCompletionRate: body.ListCompletionRate, Messages: body.Messages}
	result := initializers.DB.Create(&todolist)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create TO-DO List",
		})
		return
	}

	// when a user creates a TO-DO list, users "HasList" attribute is equal to that TO-DO lists ID
	owner.HasList = int(todolist.ID)
	// update database

	result2 := initializers.DB.Save(owner)

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update owner's TO-DO list",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "TO-DO list created successfully",
	})

}

func GetToDoLists(c *gin.Context) {

	var todolists []models.ToDoList

	// gets the current user's name,
	owner := getCurrentUser(c)

	// if list is empty;
	if owner.HasList == 0 {
		c.JSON(http.StatusOK, gin.H{
			fmt.Sprintf("TO-DO List for User %s", owner.Name): "No TO-DO list yet.",
		})
		return
	}
	// If the user is of type "second", retrieve all the ToDoLists
	if owner.Type == "second" {
		result := initializers.DB.Model(&models.ToDoList{}).Preload("Messages").Find(&todolists)
		if result.Error != nil {
			fmt.Print("hata")
		}
	} else {
		// If the user is of type "second", retrieve all the ToDoLists
		result := initializers.DB.Model(&models.ToDoList{}).Where("owner = ?", owner.Name).Preload("Messages").Find(&todolists)
		if result.Error != nil {
			fmt.Print("hata2")
		}
	}
	var listResponses []models.ToDoListResponse

	// map the messages from TO-DO lists to response entities
	for _, todolist := range todolists {
		var listResponse models.ToDoListResponse
		listResponse.Owner = todolist.Owner
		for _, message := range todolist.Messages {
			listResponse.Messages = append(listResponse.Messages, models.MessageResponse{
				Content:  message.Content,
				IsItDone: message.IsItDone,
			})
		}
		listResponse.ListCompletionRate = todolist.ListCompletionRate
		listResponses = append(listResponses, listResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		fmt.Sprintf("TO-DO List(s) for User %s", owner.Name): listResponses,
	})
}

func DeleteToDoList(c *gin.Context) {

	var todolists models.ToDoList
	var messages []models.ToDoListMessage

	owner := getCurrentUser(c)

	todolistID := owner.HasList

	// update modify time, update delete time, but not actually delete the record
	result := initializers.DB.Model(&todolists).Where("id = ?", todolistID).Update("deleted_at", time.Now())

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete TO-DO list",
		})
		return
	}

	// same for messages in the TO-DO list
	result_message := initializers.DB.Model(&messages).Where("to_do_list_id = ?", todolistID).Update("deleted_at", time.Now())

	if result_message.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "List is already empty",
		})
		return
	}

	// update user's HasList attribute after deletion
	owner.HasList = 0

	// update database for owner of the list
	result2 := initializers.DB.Save(owner)

	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update owner's TO-DO list",
		})
		return
	}

	// success message
	c.JSON(200, gin.H{
		"message": "TO-DO List deleted successfully",
	})
}

// create message for current user's TO-DO list
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

	owner := getCurrentUser(c)

	todolistID := owner.HasList

	if todolistID == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Please create a TO-DO list")
		return
	}

	message := models.ToDoListMessage{ToDoListID: todolistID, Content: body.Content, IsItDone: body.IsItDone}
	result := initializers.DB.Create(&message)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create TO-DO List",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "TO-DO message created successfully",
	})

}

// "delete" message
func DeleteToDoMessage(c *gin.Context) {

	var message models.ToDoListMessage

	messageID := c.Param("id")

	// update modify time, update delete time, but not actually delete the message
	initializers.DB.Model(&message).Where("id = ?", messageID).Update("deleted_at", time.Now())

	// success message
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Message with ID {%s} deleted", messageID),
	})
}

// update message
func UpdateToDoMessage(c *gin.Context) {

	var body struct {
		Content  string `json:"content"`
		IsItDone bool   `json:"isitdone"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read TO-DO message body",
		})
		return
	}

	var message models.ToDoListMessage

	messageID := c.Param("id")

	// users can update both content and completion
	result := initializers.DB.Model(&message).Where("id = ?", messageID).Updates(models.ToDoListMessage{Content: body.Content, IsItDone: body.IsItDone})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update TO-DO message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "TO-DO message updated successfully",
	})
}
