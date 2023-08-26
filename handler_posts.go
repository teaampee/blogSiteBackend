package main

import (
	"fmt"
	"time"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreatePost(c *gin.Context) {
	// get blog id off url
	blogIDStr := c.Param("blogID")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse blog: %v", err),
		})
	}
	// get user off req
	userIDAny, ok := c.Get("userID")
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
		Title   string
		Content string
	}
	err = c.Bind(&params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
		})
		return
	}

	// Perform input validation

	if params.Title == "" || params.Content == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}

	post, err := apiCfg.DB.CreateBlogPost(c.Request.Context(), database.CreateBlogPostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     params.Title,
		Content:   params.Content,
		UserID:    userID,
		BlogID:    blogID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create post: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func (apiCfg *apiConfig) handlerGetBlogPosts(c *gin.Context) {
	blogIDStr := c.Param("blogID")
	blogID, err := uuid.Parse(blogIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse blog's post: %v", err),
		})
	}

	posts, err := apiCfg.DB.GetBlogPosts(c.Request.Context(), blogID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch posts: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (apiCfg *apiConfig) handlerGetPost(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse post: %v", err),
		})
	}

	post, err := apiCfg.DB.GetPost(c.Request.Context(), postID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch post: %v", err),
		})
		return
	}

	// return active blogs
	c.JSON(200, gin.H{
		"post": post,
	})
}

func (apiCfg *apiConfig) handlerUpdateUserPost(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse post: %v", err),
		})
	}
	// get user off req
	userIDAny, ok := c.Get("userID")
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
		Title   string
		Content string
	}
	err = c.Bind(&params)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse json paramaters: %v", err),
		})
		return
	}

	// Perform input validation

	if params.Title == "" || params.Content == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}

	post, err := apiCfg.DB.UpdateUserPost(c.Request.Context(), database.UpdateUserPostParams{
		Title:   params.Title,
		Content: params.Content,
		ID:      postID,
		UserID:  userID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to update post: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func (apiCfg *apiConfig) handlerDeleteUserPost(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse url: %v", err),
		})
	}
	// get user off req
	userIDAny, ok := c.Get("userID")
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

	err = apiCfg.DB.DeleteUserPost(c.Request.Context(), database.DeleteUserPostParams{
		ID:     postID,
		UserID: userID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to delete post: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{})
}
