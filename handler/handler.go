package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
)

type Handler struct {
	userStore  domain.AuthStore
	actorStore domain.ActorStore
}

func NewHandler(
	us domain.AuthStore,
	as domain.ActorStore,
) *Handler {
	return &Handler{
		userStore:  us,
		actorStore: as,
	}
}
