package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var req domain.SignUpInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	err := req.HashPassword()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	res, err := h.userStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req domain.SignInInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.userStore.GetOne(strings.ToLower(req.Email), "")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	ok, err := req.CheckPassword(res.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "passwords don't match"})
	}

	accessToken, err := h.userStore.SetAccessToken(res.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	refreshToken, err := h.userStore.SetRefreshToken(res.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var accessTokenMaxAgeString = os.Getenv("ACCESS_TOKEN_MAXAGE")
	var refreshTokenMaxAgeString = os.Getenv("REFRESH_TOKEN_MAXAGE")
	var accessTokenMaxAge, _ = strconv.Atoi(accessTokenMaxAgeString)
	var refreshTokenMaxAge, _ = strconv.Atoi(refreshTokenMaxAgeString)

	c.Cookie(
		&fiber.Cookie{
			Name:     "access_token",
			Value:    *accessToken,
			Path:     "/",
			MaxAge:   accessTokenMaxAge * 60,
			Secure:   false,
			HTTPOnly: true,
		},
	)

	c.Cookie(
		&fiber.Cookie{
			Name:     "refresh_token",
			Value:    *refreshToken,
			Path:     "/",
			MaxAge:   refreshTokenMaxAge * 60,
			Secure:   false,
			HTTPOnly: true,
		},
	)

	c.Cookie(
		&fiber.Cookie{
			Name:     "logged_in",
			Value:    "true",
			Path:     "/",
			MaxAge:   accessTokenMaxAge * 60,
			Secure:   false,
			HTTPOnly: false,
		},
	)

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) RefreshAccessToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{
				"error": "no refresh token in cookies",
			},
		)
	}

	userID, err := h.userStore.GetByRefreshTokenRedis(refreshToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(err.Error())
	}

	res, err := h.userStore.GetOne("", userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	accessToken, err := h.userStore.SetAccessToken(res.ID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var accessTokenMaxAgeString = os.Getenv("ACCESS_TOKEN_MAXAGE")
	var accessTokenMaxAge, _ = strconv.Atoi(accessTokenMaxAgeString)

	c.Cookie(
		&fiber.Cookie{
			Name:     "access_token",
			Value:    *accessToken,
			Path:     "/",
			MaxAge:   accessTokenMaxAge * 60,
			Secure:   false,
			HTTPOnly: true,
			Domain:   "localhost",
		},
	)

	c.Cookie(
		&fiber.Cookie{
			Name:     "logged_in",
			Value:    "true",
			Path:     "/",
			MaxAge:   accessTokenMaxAge * 60,
			Secure:   false,
			HTTPOnly: false,
			Domain:   "localhost",
		},
	)

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"access_token": accessToken,
		},
	)
}

func (h *Handler) DeserializeUser(c *fiber.Ctx) error {
	var accessToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		accessToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		accessToken = c.Cookies("access_token")
	}

	if accessToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(
			fiber.Map{"message": "You are not logged in"},
		)
	}

	userID, tokenUUID, err := h.userStore.GetByAccessTokenRedis(accessToken)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(err.Error())
	}

	res, err := h.userStore.GetOne("", userID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	c.Locals("user", res)
	c.Locals("access_token_uuid", tokenUUID)

	return c.Next()
}

func (h *Handler) GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)

	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handler) LogoutUser(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "No refresh token in the cookies"})
	}
	accessToken := c.Locals("access_token_uuid").(string)

	err := h.userStore.DeleteTokensRedis(refreshToken, accessToken)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	now := time.Now()

	c.Cookie(
		&fiber.Cookie{
			Name:    "access_token",
			Value:   "",
			Expires: now,
		},
	)
	c.Cookie(
		&fiber.Cookie{
			Name:    "refresh_token",
			Value:   "",
			Expires: now,
		},
	)
	c.Cookie(
		&fiber.Cookie{
			Name:    "logged_in",
			Value:   "",
			Expires: now,
		},
	)

	return c.SendStatus(http.StatusOK)
}
