package main

import (
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"time"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiConfig) handlerCreateUser(c *gin.Context) {
	// getting data off request
	var params struct {
		Name     string
		Email    string
		Password string
	}
	err := c.Bind(&params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
		})
		return
	}

	// Perform input validation

	if params.Name == "" || params.Email == "" || params.Password == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}
	_, err = mail.ParseAddress(params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("false email formatting: %v", err),
		})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to hash password: %v", err),
		})
		c.Abort()
		return
	}

	// create user

	user, err := apiCfg.DB.CreateUser(c.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  string(hash),
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create user: %v", err),
		})
		return
	}

	//generate token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create token: %v", err),
		})
		return
	}
	// returning a token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	//returning it
	c.JSON(200, gin.H{
		"user": user,
	})
}

func (apiCfg *apiConfig) handlerLogin(c *gin.Context) {
	// getting data off request
	var params struct {
		Email    string
		Password string
	}
	err := c.Bind(&params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
		})
		return
	}

	// Perform input validation

	if params.Email == "" || params.Password == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}
	_, err = mail.ParseAddress(params.Email)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("false email formatting: %v", err),
		})
		return
	}

	// get user

	user, err := apiCfg.DB.GetUserByEmail(c.Request.Context(), params.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("failed to get user: %v", err),
		})
		return
	}
	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"message":   fmt.Sprintf("Invalid email or password: %v", err),
			"hpassword": user.Password,
		})
		return
	}

	//generate token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create token: %v", err),
		})
		return
	}
	// returning it
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(200, gin.H{})
}

func (apiCfg *apiConfig) handlerGetUser(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse blog: %v", err),
		})
	}

	user, err := apiCfg.DB.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch user: %v", err),
		})
	}

	c.JSON(200, gin.H{
		"message": user,
	})
}

func (apiCfg *apiConfig) handlerValidate(c *gin.Context) {
	userIDAny, _ := c.Get("userID")
	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprintln("failed to parse uuid"),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": userID,
	})
}
