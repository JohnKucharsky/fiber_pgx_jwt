package handler

import (
	"fmt"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateAddress(c *fiber.Ctx) error {
	var req domain.AddressInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}
	fmt.Println(req.Address)

	res, err := h.addressStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) GetAddresses(c *fiber.Ctx) error {
	res, err := h.addressStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOneAddress(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.addressStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateAddress(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.AddressInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.addressStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) DeleteAddress(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.addressStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}
