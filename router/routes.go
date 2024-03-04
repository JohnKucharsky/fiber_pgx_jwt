package router

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/handler"
	"github.com/JohnKucharsky/fiber_pgx_jwt/store"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Register(r *fiber.App, db *pgxpool.Pool, redis *redis.Client) {
	us := store.NewUserStore(db, redis)
	as := store.NewActorStore(db)
	cityStore := store.NewCityStore(db)
	cs := store.NewCountryStore(db)
	h := handler.NewHandler(us, as, cs, cityStore)

	v1 := r.Group("/api")

	// auth
	auth := v1.Group("/auth")
	auth.Post("/sign-up", h.SignUp)
	auth.Post("/login", h.SignIn)
	auth.Get("/logout", h.DeserializeUser, h.LogoutUser)
	auth.Get("/refresh", h.RefreshAccessToken)
	auth.Get("/me", h.DeserializeUser, h.GetMe)
	// end auth

	// actor
	actor := v1.Group("/actor")
	actor.Post("/", h.CreateActor)
	actor.Get("/", h.GetActors)
	actor.Get("/:id", h.GetOneActor)
	actor.Put("/:id", h.UpdateActor)
	actor.Delete("/:id", h.DeleteActor)
	// end actor

	// country
	country := v1.Group("/country")
	country.Post("/", h.CreateCountry)
	country.Get("/", h.GetCountries)
	country.Get("/:id", h.GetOneCountry)
	country.Put("/:id", h.UpdateCountry)
	country.Delete("/:id", h.DeleteCountry)
	// end country

	//city
	city := v1.Group("/city")
	city.Post("/", h.CreateCity)
	city.Get("/", h.GetCities)
	city.Get("/:id", h.GetOneCity)
	city.Put("/:id", h.UpdateCity)
	city.Delete("/:id", h.DeleteCity)
	// end city
}
