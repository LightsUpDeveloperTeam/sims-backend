package industry_partner

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"sims-backend/internal/utils"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}
func (h *Handler) CreateIndustryPartner(c *fiber.Ctx) error {
	var industryPartner IndustryPartner
	if err := c.BodyParser(&industryPartner); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed industry partner payload: %+v", industryPartner)

	if err := h.Service.CreateIndustryPartner(&industryPartner); err != nil {
		log.Printf("Error creating industry partner: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create industry partner",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"industry partner created successfully",
		industryPartner,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetIndustryPartner(c *fiber.Ctx) error {
	industryPartner, err := h.Service.GetIndustryPartner()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get industry partner",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Industry partner retrieved successfully",
		industryPartner,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetIndustryPartnerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid industry partner ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	industryPartner, err := h.Service.GetIndustryPartnerByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get industry partner",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Industry partner retrieved successfully",
		industryPartner,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateIndustryPartner(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid industry partner ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var industryPartner IndustryPartner
	if err := c.BodyParser(&industryPartner); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	industryPartner.ID = uint64(id)

	if err := h.Service.UpdateIndustryPartner(&industryPartner); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Industry partner updated successfully",
		industryPartner,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteIndustryPartner(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting user ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteIndustryPartner(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting user by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "industry partner deleted successfully", nil, nil, nil, nil, nil,
	))
}
