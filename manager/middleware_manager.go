package manager

import (
	"itdp-group3-backend/auth"
	"itdp-group3-backend/middleware"
)

type MiddlewareManager interface {
	AuthMiddleware() middleware.AuthTokenMiddleware
}

type middlewareManager struct {
	auth auth.Token
}

func NewMiddlewareManager(auth auth.Token) MiddlewareManager {
	return &middlewareManager{
		auth: auth,
	}
}

func (m *middlewareManager) AuthMiddleware() middleware.AuthTokenMiddleware {
	return middleware.NewTokenValidator(m.auth)
}
