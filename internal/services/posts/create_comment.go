package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, userID, postID int64, req posts.CreateCommentRequest) error {
	existingPost, err := s.postRepo.GetPostByID(ctx, postID, userID)

	if err != nil {
		return err
	}

	if existingPost == nil {
		return errors.New("post tersebut tidak ditemukan")
	}

	now := time.Now()

	model := posts.CommentModel{
		UserID:         userID,
		PostID:         postID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}

	err = s.postRepo.CreateComment(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("error while creating a comment")
		return err
	}

	return nil
}
