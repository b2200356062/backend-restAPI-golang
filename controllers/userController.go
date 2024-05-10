package controllers

import (
	"fmt"
	"net/http"
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
		Type     string
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

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash), Type: body.Type}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user: user already exists in database",
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
		Type     string
	}

	if c.Bind(&body) != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	var user models.User

	initializers.DB.First(&user, "email = ?", body.Email)

	// if no user with the requested email
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

	// get random secret token and hash it
	tokenString, err := token.SignedString([]byte("sgljg3ı2jg902gfjdskfasjfndlFSOEJFSAFO2IFVUNDWşdfmhbng230f9j2efndsjgna423u"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create JWT token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	// login using cookies to send and get jwt token
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Login successfull, JWT Token for this user: %s", tokenString),
	})
}

// gets a pointer to current logged in user
func getCurrentUser(c *gin.Context) *models.User {

	user, _ := c.Get("user")

	userModel, ok := user.(models.User)
	if !ok {

		return nil
	}
	return &userModel
}
