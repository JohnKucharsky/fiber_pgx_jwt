package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateCity(c *fiber.Ctx) error {
	var req domain.CityInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	id, err := h.cityStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.cityStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) GetCities(c *fiber.Ctx) error {
	res, err := h.cityStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOneCity(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.cityStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateCity(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.CityInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	id, err = h.cityStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	res, err := h.cityStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) DeleteCity(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	id, err = h.cityStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(id)
}
