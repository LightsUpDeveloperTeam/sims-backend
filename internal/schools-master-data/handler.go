package schoolsmasterdata

import (
	"log"
	"sims-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}


func (h *Handler) CreateSchools(c *fiber.Ctx) error {
	var school School
	if err := c.BodyParser(&school); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed school payload: %+v", school)

	if err := h.Service.CreateSchool(&school); err != nil {
		log.Printf("Error creating school: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create school",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"School created successfully",
		school,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetSchools(c *fiber.Ctx) error {
	schools, err := h.Service.GetSchools()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get schools",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Schools retrieved successfully",
		schools,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetSchoolByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid school ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	school, err := h.Service.GetSchoolByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get school",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"School retrieved successfully",
		school,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateSchool(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid school ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var school School
	if err := c.BodyParser(&school); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	school.ID = uint64(id)

	if err := h.Service.UpdateSchool(&school); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"School updated successfully",
		school,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteSchool(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid school ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	if err := h.Service.DeleteSchool(uint64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"School deleted successfully",
		nil,
		nil, nil, nil, nil,
	))
}
