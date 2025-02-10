package posts

import (
	"context"

	"github.com/amiulam/simple-forum/internal/middleware"
	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	GetAllPost(ctx context.Context, userID int64, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
	GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error)

	CreateComment(ctx context.Context, userID, postID int64, req posts.CreateCommentRequest) error

	UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error
}

type Handler struct {
	*gin.Engine
	postSvc postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())
	route.GET("/", h.GetAllPost)
	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
	route.GET("/:postID", h.GetPostByID)
}
