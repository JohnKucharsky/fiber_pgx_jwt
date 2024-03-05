package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
)

type Handler struct {
	userStore     domain.AuthStore
	actorStore    domain.ActorStore
	countryStore  domain.CountryStore
	cityStore     domain.CityStore
	addressStore  domain.AddressStore
	categoryStore domain.CategoryStore
	languageStore domain.LanguageStore
}

func NewHandler(
	us domain.AuthStore,
	as domain.ActorStore,
	cs domain.CountryStore,
	cityStore domain.CityStore,
	addressStore domain.AddressStore,
	categoryStore domain.CategoryStore,
	languageStore domain.LanguageStore,
) *Handler {
	return &Handler{
		userStore:     us,
		actorStore:    as,
		countryStore:  cs,
		cityStore:     cityStore,
		addressStore:  addressStore,
		categoryStore: categoryStore,
		languageStore: languageStore,
	}
}
