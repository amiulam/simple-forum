package posts

import (
	"context"
	"database/sql"
	"strings"

	"github.com/amiulam/simple-forum/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by) values (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPostByID(ctx context.Context, postID, userID int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, COALESCE(uv.is_liked, false) as is_liked FROM posts p JOIN users u ON u.id = p.user_id LEFT JOIN user_activities uv ON uv.post_id = p.id AND uv.user_id = ? WHERE p.id = ?`

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)

	row := r.db.QueryRowContext(ctx, query, userID, postID)

	err := row.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags, &isLiked)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &posts.Post{
		ID:           model.ID,
		UserID:       model.UserID,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLiked,
	}, nil
}

func (r *repository) GetAllPost(ctx context.Context, userID int64, limit, offset int) (posts.GetAllPostResponse, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, COALESCE(ua.is_liked, false) FROM posts p JOIN users u ON u.id = p.user_id LEFT JOIN user_activities ua ON ua.post_id = p.id AND ua.user_id = ? ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	response := posts.GetAllPostResponse{}

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)

	if err != nil {
		return response, err
	}

	defer rows.Close()

	data := make([]posts.Post, 0)

	for rows.Next() {
		var (
			model    posts.PostModel
			username string
			isLiked  bool
		)

		err = rows.Scan(&model.ID, &model.UserID, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags, &isLiked)

		if err != nil {
			return response, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserID:       model.UserID,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
			IsLiked:      isLiked,
		})
	}

	postCount := r.GetPostCount(ctx)

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  limit,
		Offset: offset,
	}
	response.Count = postCount
	return response, nil
}

func (r *repository) GetPostCount(ctx context.Context) int64 {
	query := `SELECT COUNT(*) as count FROM posts`

	var count int64

	row := r.db.QueryRowContext(ctx, query)

	row.Scan(&count)

	return count
}
