package main

import (
	"fmt"
	"time"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateBlog(c *gin.Context) {
	// get user off req
	userIDAny, ok := c.Get("UserID")
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprint("failed to get the user id off req"),
		})
		return
	}
	//convert into uuid
	userID, ok := userIDAny.(uuid.UUID)
	if !ok {
		c.JSON(400, gin.H{
			"message": fmt.Sprint("failed to parse uuid"),
		})
		return
	}
	// getting data off request
	var params struct {
		Title       string
		Description string
	}
	err := c.Bind(&params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
		})
		return
	}

	// Perform input validation

	if params.Title == "" || params.Description == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}

	blog, err := apiCfg.DB.CreateBlog(c.Request.Context(), database.CreateBlogParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       params.Title,
		Description: params.Description,
		UserID:      userID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create blog: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"blog": blog,
	})
}

// func (apiCfg *apiConfig) handlerLogin(c *gin.Context) {
// 	// getting data off request
// 	var params struct {
// 		Email    string
// 		Password string
// 	}
// 	err := c.Bind(&params)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
// 		})
// 		return
// 	}

// 	// Perform input validation

// 	if params.Email == "" || params.Password == "" {
// 		c.JSON(400, gin.H{
// 			"message": fmt.Sprintf("missing required fields: %v", err),
// 		})
// 		return
// 	}
// 	_, err = mail.ParseAddress(params.Email)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"message": fmt.Sprintf("false email formatting: %v", err),
// 		})
// 		return
// 	}

// 	// get user

// 	user, err := apiCfg.DB.GetUserByEmail(c.Request.Context(), params.Email)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": fmt.Sprintf("failed to get user: %v", err),
// 		})
// 		return
// 	}
// 	// validate password
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"message":   fmt.Sprintf("Invalid email or password: %v", err),
// 			"hpassword": user.Password,
// 		})
// 		return
// 	}

// 	//generate token
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
// 		"sub": user.ID,
// 		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"message": fmt.Sprintf("failed to create token: %v", err),
// 		})
// 		return
// 	}
// 	// returning it
// 	c.SetSameSite(http.SameSiteLaxMode)
// 	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
// 	c.JSON(200, gin.H{})
// }

// func (apiCfg *apiConfig) handlerValidate(c *gin.Context) {
// 	user, _ := c.Get("user")

// 	c.JSON(200, gin.H{
// 		"message": user,
// 	})
// }
