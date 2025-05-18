package controllers

import (
	"go-mvc-demo/config"
	"go-mvc-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func NewCommentController() *CommentController {
	return &CommentController{}
}

func (ctl *CommentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&comment)
	c.JSON(http.StatusOK, comment)
}

func (ctl *CommentController) UpdateComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	var updatedData models.Comment
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.Content = updatedData.Content
	comment.PostID = updatedData.PostID

	if err := config.DB.Save(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully", "comment": comment})
}

func (ctl *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	if err := config.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
