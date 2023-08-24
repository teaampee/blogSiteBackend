package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) UserAuth(c *gin.Context) {
	//get cookie
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode and validate
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("token expired: %v", err),
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//get user
		//convert subject to uuid
		userID, _ := uuid.Parse(claims["sub"].(string))

		user, err := apiCfg.DB.GetUserByID(c.Request.Context(), userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("failed to get user: %v", err),
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//attach to req
		c.Set("user", user)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("invalid token: %v", err),
		})
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//continu
	c.Next()
}
