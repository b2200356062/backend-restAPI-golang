package controllers

import (
	"fmt"
	"net/http"
	"os"
	"staj/initializers"
	"staj/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// signup function
func SignUp(c *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	// hashing password

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed generate hash from password",
		})
		return
	}

	fmt.Print("after hash ")

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	fmt.Print("after user ")

	result := initializers.DB.Create(&user)

	fmt.Print("after result")

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user created successfully",
	})
}

func Login(c *gin.Context) {

	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})

		return
	}

	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	// if no user with the requested email in database
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})

		return
	}

	// compare the password in request and in database
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid password",
		})

		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), // expires in 30 days
	})

	// get token from environment and hash it
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create JWT token",
		})

		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})

}
