package BKK

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

// Internship Vacancy Handler

func (h *Handler) CreateVacancy(c *fiber.Ctx) error {
	var internshipVacancy InternshipVacancy
	if err := c.BodyParser(&internshipVacancy); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload",
			nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Parsed vacancy payload: %+v", internshipVacancy)

	if err := h.Service.CreateVacancy(&internshipVacancy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(),
			nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "Vacancy created successfully", internshipVacancy,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetVacancies(c *fiber.Ctx) error {
	vacancies, err := h.Service.GetVacancies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get vacancies",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Vacancies retrieved successfully", vacancies,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetVacancyByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid Vacancy ID",
			nil, nil, nil, nil, nil,
		))
	}

	vacancy, err := h.Service.GetVacancyByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get vacancy",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Vacancy retrieved successfully", vacancy,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateVacancy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid vacancy ID",
			nil, nil, nil, nil, nil,
		))
	}

	var vacancy InternshipVacancy
	if err := c.BodyParser(&vacancy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil, nil, nil, nil, nil,
		))
	}
	vacancy.ID = uint64(id)

	if err := h.Service.UpdateVacancy(&vacancy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil, nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Vacancy updated successfully", vacancy,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteVacancy(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting user ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteVacancy(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting user by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "internship vacancy deleted successfully", nil, nil, nil, nil, nil,
	))
}

// Internship Registration Handler

func (h *Handler) CreateRegistration(c *fiber.Ctx) error {
	var internshipRegistration InternshipRegistration
	if err := c.BodyParser(&internshipRegistration); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload",
			nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Parsed registration payload: %+v", internshipRegistration)

	if err := h.Service.CreateRegistration(&internshipRegistration); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(),
			nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "Registration created successfully", internshipRegistration,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetRegistration(c *fiber.Ctx) error {
	registration, err := h.Service.GetRegistrations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get registrations",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Registrations retrieved successfully", registration,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetRegistrationByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid registration ID",
			nil, nil, nil, nil, nil,
		))
	}

	registration, err := h.Service.GetRegistrationByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get registration",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Registration retrieved successfully", registration,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteRegistration(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting internship registration ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteRegistration(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting internship registration by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Internship registration deleted successfully", nil, nil, nil, nil, nil,
	))
}

// Alumnus Distribution Handler

func (h *Handler) CreateDistribution(c *fiber.Ctx) error {
	var alumnusDistribution AlumnusDistribution
	if err := c.BodyParser(&alumnusDistribution); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload",
			nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Parsed vacancy payload: %+v", alumnusDistribution)

	if err := h.Service.CreateDistribution(&alumnusDistribution); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(),
			nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "Alumnus distribution created successfully", alumnusDistribution,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetDistribution(c *fiber.Ctx) error {
	distribution, err := h.Service.GetDistribution()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get distribution",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Alumnus distribution retrieved successfully", distribution,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetDistributionByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid alumnus distribution ID",
			nil, nil, nil, nil, nil,
		))
	}

	alumnusDistribution, err := h.Service.GetDistributionByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get alumnus distribution",
			nil, nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Alumnus Distribution retrieved successfully", alumnusDistribution,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateDistribution(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting alumnus distribution ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid alumnus distribution ID", nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Received request to update alumnus distribution: %d", id)
	var alumnusDistribution AlumnusDistribution
	if err := c.BodyParser(&alumnusDistribution); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	alumnusDistribution.ID = uint64(id)
	if err := h.Service.UpdateDistribution(&alumnusDistribution); err != nil {
		log.Printf("Error updating alumnus distribution: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Updated alumnus distribution: %+v", alumnusDistribution)
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Alumnus distribution updated successfully", alumnusDistribution, nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteDistribution(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting alumnus distrtibution ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid alumnus distribution ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteDistribution(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting alumnus distribution by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Alumnus distribution deleted successfully", nil, nil, nil, nil, nil,
	))
}
