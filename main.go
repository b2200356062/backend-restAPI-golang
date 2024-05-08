package main

import (
	"staj/controllers"
	"staj/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	router.Run()

}
