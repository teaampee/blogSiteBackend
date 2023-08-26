package main

import (
	"fmt"
	"time"

	"example.com/blog/v2/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateComment(c *gin.Context) {
	postIDStr := c.Query("post_id")
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

	if params.Content == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}

	comment, err := apiCfg.DB.CreateComment(c.Request.Context(), database.CreateCommentParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Content:   params.Content,
		UserID:    userID,
		PostID:    postID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create comment: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"comment": comment,
	})
}

func (apiCfg *apiConfig) handlerGetPostComments(c *gin.Context) {
	postIDStr := c.Param("postID")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse post's url: %v", err),
		})
	}

	comments, err := apiCfg.DB.GetPostComments(c.Request.Context(), postID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch post's comments: %v", err),
		})
		return
	}
	c.JSON(200, gin.H{
		"comments": comments,
	})
}

func (apiCfg *apiConfig) handlerGetComment(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse comment: %v", err),
		})
	}

	comment, err := apiCfg.DB.GetComment(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to fetch comment: %v", err),
		})
		return
	}

	// return comment
	c.JSON(200, gin.H{
		"comment": comment,
	})
}

func (apiCfg *apiConfig) handlerUpdateComment(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	commentID, err := uuid.Parse(commentIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to parse comment: %v", err),
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

	if params.Content == "" {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("missing required fields: %v", err),
		})
		return
	}

	comment, err := apiCfg.DB.UpdateUserComment(c.Request.Context(), database.UpdateUserCommentParams{
		Content: params.Content,
		UserID:  userID,
		ID:      commentID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to create comment: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{
		"comment": comment,
	})
}

func (apiCfg *apiConfig) handlerDeleteUserComment(c *gin.Context) {
	commentIDStr := c.Param("commentID")
	commentID, err := uuid.Parse(commentIDStr)
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

	err = apiCfg.DB.DeleteUserComment(c.Request.Context(), database.DeleteUserCommentParams{
		ID:     commentID,
		UserID: userID,
	})

	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("failed to delete comment: %v", err),
		})
		return
	}

	//returning it
	c.JSON(200, gin.H{})
}
