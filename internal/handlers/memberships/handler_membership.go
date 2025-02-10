package memberships

import (
	"context"

	"github.com/amiulam/simple-forum/internal/middleware"
	"github.com/amiulam/simple-forum/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error)
	Logout(ctx context.Context, userID int64) error
}

type Handler struct {
	*gin.Engine
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeProtected := h.Group("memberships")
	routeProtected.Use(middleware.AuthMiddleware())
	routeProtected.POST("/logout", h.Logout)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthMiddlewareForRefreshToken())
	routeRefresh.POST("/refresh-token", h.RefreshToken)
}
