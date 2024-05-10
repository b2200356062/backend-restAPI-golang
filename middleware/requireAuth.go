package middleware

import (
	"fmt"
	"staj/initializers"
	"staj/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// get cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		// http status unauthorized
		c.AbortWithStatus(401)
	}

	// validate
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// secret for jwt token
		return []byte("sgljg3ı2jg902gfjdskfasjfndlFSOEJFSAFO2IFVUNDWşdfmhbng230f9j2efndsjgna423u"), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(401, "Error: You are not authorized. Please Login.")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(401, "Error: JWT Token is expired.")
		}

		// find the user with the correct token
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		// if there is not, that means no one has logged in
		if user.ID == 0 {
			c.AbortWithStatusJSON(401, "Error: You are not authorized. Please Login.")
		}

		// attach to request
		c.Set("user", user)

		// continue
		c.Next()

	} else {
		c.AbortWithStatus(401)
	}

}
