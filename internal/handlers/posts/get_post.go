package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPostByID(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetInt64("userID")
	postID, err := strconv.ParseInt(c.Param("postID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid post id").Error(),
		})
		return
	}

	response, err := h.postSvc.GetPostByID(ctx, postID, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
