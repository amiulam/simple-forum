package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpsertUserActivity(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.UserActivityRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")
	postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("post id pada param tidak valid").Error(),
		})
		return
	}

	err = h.postSvc.UpsertUserActivity(ctx, postID, userID, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user activity saved",
	})
}
