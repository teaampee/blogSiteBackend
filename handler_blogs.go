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

func (apiCfg *apiConfig) handlerGetBlogs(c *gin.Context) {

	blogs, err := apiCfg.DB.GetBlogs(c.Request.Context())
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch blogs: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"blogs": blogs,
	})
}

func (apiCfg *apiConfig) handlerGetBlog(c *gin.Context) {
	blogIDStr := c.Param("blogID")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse blog: %v", err),
		})
	}

	blog, err := apiCfg.DB.GetBlog(c.Request.Context(), blogID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch blog: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"blog": blog,
	})
}

func (apiCfg *apiConfig) handlerGetUserBlog(c *gin.Context) {
	userIDStr := c.Param("userID")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse blog: %v", err),
		})
	}

	blog, err := apiCfg.DB.GetUserBlog(c.Request.Context(), userID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch blog: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"blog": blog,
	})
}

func (apiCfg *apiConfig) handlerUpdateUserBlog(c *gin.Context) {
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

	blog, err := apiCfg.DB.UpdateUserBlog(c.Request.Context(), database.UpdateUserBlogParams{
		Title:       params.Title,
		Description: params.Description,
		UserID:      userID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to update blog: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"blog": blog,
	})
}

func (apiCfg *apiConfig) handlerDeleteUserBlog(c *gin.Context) {
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

	err := apiCfg.DB.DeleteUserBlog(c.Request.Context(), userID)

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create blog: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{})
}
