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
	addressStore := store.NewAddressStore(db)
	cs := store.NewCountryStore(db)
	categoryStore := store.NewCategoryStore(db)
	languageStore := store.NewLanguageStore(db)
	customerStore := store.NewCustomerStore(db)
	staffStore := store.NewStaffStore(db)

	h := handler.NewHandler(
		us,
		as,
		cs,
		cityStore,
		addressStore,
		categoryStore,
		languageStore,
		customerStore,
		staffStore,
	)

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
	actor.Post("/", h.DeserializeUser, h.CreateActor)
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

	//address
	address := v1.Group("/address")
	address.Post("/", h.CreateAddress)
	address.Get("/", h.GetAddresses)
	address.Get("/:id", h.GetOneAddress)
	address.Put("/:id", h.UpdateAddress)
	address.Delete("/:id", h.DeleteAddress)
	// end address

	// category
	category := v1.Group("/category")
	category.Post("/", h.CreateCategory)
	category.Get("/", h.GetCategories)
	category.Get("/:id", h.GetOneCategory)
	category.Put("/:id", h.UpdateCategory)
	category.Delete("/:id", h.DeleteCategory)
	// end category

	// language
	language := v1.Group("/language")
	language.Post("/", h.CreateLanguage)
	language.Get("/", h.GetLanguages)
	language.Get("/:id", h.GetOneLanguage)
	language.Put("/:id", h.UpdateLanguage)
	language.Delete("/:id", h.DeleteLanguage)
	// end language

	// customer
	customer := v1.Group("/customer")
	customer.Post("/", h.CreateCustomer)
	customer.Get("/", h.GetCustomers)
	customer.Get("/:id", h.GetOneCustomer)
	customer.Put("/:id", h.UpdateCustomer)
	customer.Delete("/:id", h.DeleteCustomer)
	// end customer

	// staff
	staff := v1.Group("/staff")
	staff.Post("/", h.CreateStaff)
	staff.Get("/", h.GetStaffs)
	staff.Get("/:id", h.GetOneStaff)
	staff.Put("/:id", h.UpdateStaff)
	staff.Delete("/:id", h.DeleteStaff)
	// end staff
}
