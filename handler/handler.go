package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
)

type Handler struct {
	userStore domain.AuthStore
}

func NewHandler(
	us domain.AuthStore,
) *Handler {
	return &Handler{
		userStore: us,
	}
}
