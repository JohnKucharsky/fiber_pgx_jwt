package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateStore(c *fiber.Ctx) error {
	var req domain.StoreInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.storeStore.Create(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	address, err := h.addressStore.GetOne(res.AddressID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	staff, err := h.staffStore.GetOne(res.ManagerID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var addrID int
	if staff.AddressID != nil {
		addrID = *staff.AddressID
	}
	staffAddress, err := h.addressStore.GetOne(addrID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var stf = domain.StaffDBtoStaff(staff, staffAddress)

	return c.Status(http.StatusCreated).JSON(
		domain.StoreDBtoStore(
			res,
			address,
			stf,
		),
	)
}

func (h *Handler) GetStores(c *fiber.Ctx) error {
	resStores, err := h.storeStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(domain.StoreDBSecondVerToStore(resStores))
}

func (h *Handler) GetOneStore(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.storeStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	address, err := h.addressStore.GetOne(res.AddressID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	staff, err := h.staffStore.GetOne(res.ManagerID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var addrID int
	if staff.AddressID != nil {
		addrID = *staff.AddressID
	}
	staffAddress, err := h.addressStore.GetOne(addrID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var stf = domain.StaffDBtoStaff(staff, staffAddress)

	return c.Status(http.StatusCreated).JSON(
		domain.StoreDBtoStore(
			res,
			address,
			stf,
		),
	)
}

func (h *Handler) UpdateStore(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.StoreInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.storeStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	address, err := h.addressStore.GetOne(res.AddressID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	staff, err := h.staffStore.GetOne(res.ManagerID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var addrID int
	if staff.AddressID != nil {
		addrID = *staff.AddressID
	}
	staffAddress, err := h.addressStore.GetOne(addrID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var stf = domain.StaffDBtoStaff(staff, staffAddress)

	return c.Status(http.StatusCreated).JSON(
		domain.StoreDBtoStore(
			res,
			address,
			stf,
		),
	)
}

func (h *Handler) DeleteStore(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.storeStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	address, err := h.addressStore.GetOne(res.AddressID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	staff, err := h.staffStore.GetOne(res.ManagerID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var addrID int
	if staff.AddressID != nil {
		addrID = *staff.AddressID
	}
	staffAddress, err := h.addressStore.GetOne(addrID)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	var stf = domain.StaffDBtoStaff(staff, staffAddress)

	return c.Status(http.StatusCreated).JSON(
		domain.StoreDBtoStore(
			res,
			address,
			stf,
		),
	)
}
