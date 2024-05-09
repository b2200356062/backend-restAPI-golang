package main

import (
	"staj/controllers"
	"staj/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.POST("/signup", controllers.SignUp)

	router.POST("/login", controllers.Login)

	router.POST("/createlist", controllers.CreateTODOList)

	router.GET("/getlists", controllers.GetToDoLists)

	router.PUT("/deletelist/:id", controllers.DeleteToDoList)

	router.POST("/createmessage", controllers.CreateToDoMessage)

	router.PUT("/deletemessage/:id", controllers.DeleteToDoMessage)

	router.Run("localhost:8080")

}
