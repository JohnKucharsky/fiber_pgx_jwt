package handler

import (
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/JohnKucharsky/fiber_pgx_jwt/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateCustomer(c *fiber.Ctx) error {
	var req domain.CustomerInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.customerStore.Create(req)
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
		domain.CustomerDBtoCustomer(
			res,
			addr,
		),
	)
}

func (h *Handler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.customerStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	addresses, err := h.addressStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var resCustomers []domain.Customer

	for _, customer := range customers {
		var addr *domain.Address
		if customer.AddressID != nil {
			for _, address := range addresses {
				if address.ID == *customer.AddressID {
					addr = address
				}
			}
			resCustomers = append(
				resCustomers,
				domain.CustomerDBtoCustomer(customer, addr),
			)
		}
	}

	return c.Status(http.StatusOK).JSON(resCustomers)
}

func (h *Handler) GetOneCustomer(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.customerStore.GetOne(id)
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
		domain.CustomerDBtoCustomer(
			res,
			addr,
		),
	)
}

func (h *Handler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.CustomerInput
	if err := utils.Bind(c, &req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.customerStore.Update(req, id)
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
		domain.CustomerDBtoCustomer(
			res,
			addr,
		),
	)
}

func (h *Handler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.customerStore.Delete(id)
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
		domain.CustomerDBtoCustomer(
			res,
			addr,
		),
	)
}
