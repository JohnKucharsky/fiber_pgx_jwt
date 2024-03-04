package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
)

type Handler struct {
	userStore    domain.AuthStore
	actorStore   domain.ActorStore
	countryStore domain.CountryStore
	cityStore    domain.CityStore
	addressStore domain.AddressStore
}

func NewHandler(
	us domain.AuthStore,
	as domain.ActorStore,
	cs domain.CountryStore,
	cityStore domain.CityStore,
	addressStore domain.AddressStore,
) *Handler {
	return &Handler{
		userStore:    us,
		actorStore:   as,
		countryStore: cs,
		cityStore:    cityStore,
		addressStore: addressStore,
	}
}
