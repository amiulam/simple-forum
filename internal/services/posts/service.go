package posts

import (
	"context"

	"github.com/amiulam/simple-forum/internal/configs"
	"github.com/amiulam/simple-forum/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	GetPostByID(ctx context.Context, postID, userID int64) (*posts.Post, error)
	GetAllPost(ctx context.Context, userID int64, limit, offset int) (posts.GetAllPostResponse, error)

	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetCommentsByPostID(ctx context.Context, postID int64) ([]posts.Comment, error)

	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	CountLikeByPostID(ctx context.Context, postID int64) (int, error)
}

type service struct {
	postRepo postRepository
	cfg      *configs.Config
}

func NewService(postRepo postRepository, cfg *configs.Config) *service {
	return &service{
		postRepo: postRepo,
		cfg:      cfg,
	}
}
