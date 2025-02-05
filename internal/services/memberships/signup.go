package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/amiulam/simple-forum/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username, 0)

	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("username or email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	now := time.Now()

	model := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
	}

	return s.membershipRepo.CreateUser(ctx, model)
}
