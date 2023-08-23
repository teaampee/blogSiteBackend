package main

import (
	"fmt"
	"net/mail"
	"time"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	// create user

	user, err := apiCfg.DB.CreateUser(c.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
		Password:  params.Password,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create user: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"user": user,
	})
}
