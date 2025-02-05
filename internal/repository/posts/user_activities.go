package posts

import (
	"context"
	"database/sql"

	"github.com/amiulam/simple-forum/internal/model/posts"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error) {
	query := `SELECT * from user_activities WHERE user_id=? AND post_id=?`

	var response posts.UserActivityModel

	row := r.db.QueryRowContext(ctx, query, model.UserID, model.PostID)

	err := row.Scan(&response.ID, &response.UserID, &response.PostID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `INSERT INTO user_activities (user_id, post_id, is_liked, created_at, updated_at, created_by, updated_by) values (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error {
	query := `UPDATE user_activities SET is_liked=?, updated_at=?, updated_by=? WHERE post_id=? AND user_id=?`

	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.PostID, model.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CountLikeByPostID(ctx context.Context, postID int64) (int, error) {
	query := `SELECT COUNT(id) from user_activities WHERE post_id = ? AND is_liked = true`

	var response int

	row := r.db.QueryRowContext(ctx, query, postID)

	err := row.Scan(&response)

	if err != nil {
		if err == sql.ErrNoRows {
			return response, nil
		}
		return response, err
	}

	return response, nil
}
