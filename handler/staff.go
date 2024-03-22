package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateStaff(c *fiber.Ctx) error {
	var req domain.StaffInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.staffStore.Create(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var addr *domain.Address
	if res.AddressID != nil {
		address, err := h.addressStore.GetOne(*res.AddressID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		addr = address
	}

	return c.Status(http.StatusCreated).JSON(
		domain.StaffDBtoStaff(
			res,
			addr,
		),
	)
}

func (h *Handler) GetStaffs(c *fiber.Ctx) error {
	customers, err := h.staffStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	addresses, err := h.addressStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var resStaffs []domain.Staff

	for _, stf := range customers {
		var addr *domain.Address
		if stf.AddressID != nil {
			for _, address := range addresses {
				if address.ID == *stf.AddressID {
					addr = address
				}
			}
			resStaffs = append(
				resStaffs,
				domain.StaffDBtoStaff(stf, addr),
			)
		}
	}

	return c.Status(http.StatusOK).JSON(resStaffs)
}

func (h *Handler) GetOneStaff(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.staffStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var addr *domain.Address
	if res.AddressID != nil {
		address, err := h.addressStore.GetOne(*res.AddressID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		addr = address
	}

	return c.Status(http.StatusOK).JSON(
		domain.StaffDBtoStaff(
			res,
			addr,
		),
	)
}

func (h *Handler) UpdateStaff(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.StaffInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.staffStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var addr *domain.Address
	if res.AddressID != nil {
		address, err := h.addressStore.GetOne(*res.AddressID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		addr = address
	}

	return c.Status(http.StatusCreated).JSON(
		domain.StaffDBtoStaff(
			res,
			addr,
		),
	)
}

func (h *Handler) DeleteStaff(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.staffStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var addr *domain.Address
	if res.AddressID != nil {
		address, err := h.addressStore.GetOne(*res.AddressID)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err.Error())
		}
		addr = address
	}

	return c.Status(http.StatusCreated).JSON(
		domain.StaffDBtoStaff(
			res,
			addr,
		),
	)
}
