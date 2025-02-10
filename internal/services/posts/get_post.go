package posts

import (
	"context"

	"github.com/amiulam/simple-forum/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID, userID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID, userID)

	if err != nil {
		log.Error().Err(err).Msg("error get post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)

	if err != nil {
		log.Error().Err(err).Msg("error get count like")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentsByPostID(ctx, postID)

	if err != nil {
		log.Error().Err(err).Msg("error get comments from database")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserID:       postDetail.UserID,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil

}
