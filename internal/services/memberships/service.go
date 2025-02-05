package memberships

import (
	"context"
	"time"

	"github.com/amiulam/simple-forum/internal/configs"
	"github.com/amiulam/simple-forum/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

type service struct {
	membershipRepo membershipRepository
	cfg            *configs.Config
}

func NewService(membershipRepo membershipRepository, cfg *configs.Config) *service {
	return &service{
		membershipRepo: membershipRepo,
		cfg:            cfg,
	}
}
