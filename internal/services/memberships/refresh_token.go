package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/amiulam/simple-forum/internal/model/memberships"
	"github.com/amiulam/simple-forum/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("error get refresh token from database")
		return "", err

	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has expired")
	}

	// token not match
	if existingRefreshToken.RefreshToken != req.Token {
		return "", errors.New("refresh token is invalid")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("user not exist")
	}

	accessToken, err := jwt.CreateToken(userID, user.Username, s.cfg.Service.SecretJWT)

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
