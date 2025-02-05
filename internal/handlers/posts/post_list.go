package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	pageSize, err := strconv.Atoi(c.Query("pageSize"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page size").Error(),
		})
		return
	}

	pageIndex, err := strconv.Atoi(c.Query("pageIndex"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page index").Error(),
		})
		return
	}

	response, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
