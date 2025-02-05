package posts

import (
	"net/http"

	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreatePostRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// get UserID in context that set in the middleware earlier
	userID := c.GetInt64("userID")

	err := h.postSvc.CreatePost(ctx, userID, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "post created successfully",
	})
}
