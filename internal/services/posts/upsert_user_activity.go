package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, req posts.UserActivityRequest) error {
	existingPost, err := s.postRepo.GetPostByID(ctx, postID)

	if err != nil {
		return err
	}

	if existingPost == nil {
		return errors.New("post tersebut tidak ditemukan")
	}

	now := time.Now()

	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}

	if userActivity == nil {
		if !req.IsLiked {
			return errors.New("anda belum pernah like sebelumnya")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("error create or update user activity to database")
		return err
	}

	return nil
}
