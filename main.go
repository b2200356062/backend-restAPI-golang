package main

import (
	"staj/controllers"
	"staj/initializers"
	"staj/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.POST("/signup", controllers.SignUp)

	router.POST("/login", controllers.Login)

	router.POST("/createlist", middleware.RequireAuth, controllers.CreateTODOList)

	router.GET("/getlists", middleware.RequireAuth, controllers.GetToDoLists)

	router.PUT("/deletelist", middleware.RequireAuth, controllers.DeleteToDoList)

	router.POST("/createmessage", middleware.RequireAuth, controllers.CreateToDoMessage)

	router.PUT("/deletemessage/:id", middleware.RequireAuth, controllers.DeleteToDoMessage)

	router.PUT("/updatemessage/:id", middleware.RequireAuth, controllers.UpdateToDoMessage)

	//router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run("localhost:8080")

}
