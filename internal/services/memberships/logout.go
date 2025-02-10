package memberships

import (
	"context"
)

func (s *service) Logout(ctx context.Context, userID int64) error {
	// Delete/invalidate refresh token
	return s.membershipRepo.DeleteRefreshToken(ctx, userID)
}
