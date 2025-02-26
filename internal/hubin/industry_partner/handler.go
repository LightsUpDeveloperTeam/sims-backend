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

//Collaboration History Handler

func (h *Handler) CreateCollaborationHistory(c *fiber.Ctx) error {
	var collaborationHistory CollaborationHistory
	if err := c.BodyParser(&collaborationHistory); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed collaboration history payload: %+v", collaborationHistory)

	if err := h.Service.CreateCollaborationHistory(&collaborationHistory); err != nil {
		log.Printf("Error creating collaboration history: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create collaboration history",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Collaboration created successfully",
		collaborationHistory,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetCollaborationHistory(c *fiber.Ctx) error {
	collaborationHistory, err := h.Service.GetCollaborationHistory()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get collaboration history",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Collaboration history retrieved successfully",
		collaborationHistory,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetCollaborationHistoryByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid collaboration history ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	collaborationHistory, err := h.Service.GetCollaborationHistoryByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get collaboration history",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Collaboration history retrieved successfully",
		collaborationHistory,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateCollaborationHistory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid collaboration history ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var collaborationHistory CollaborationHistory
	if err := c.BodyParser(&collaborationHistory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	collaborationHistory.ID = uint64(id)

	if err := h.Service.UpdateCollaborationHistory(&collaborationHistory); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Collaboration history updated successfully",
		collaborationHistory,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteCollaborationHistory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting collaboration history ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid collaboration history ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteCollaborationHistory(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting collaboration history by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Collaboration history deleted successfully", nil, nil, nil, nil, nil,
	))
}

//Memorandum Of Understanding Handlers

func (h *Handler) CreateMemorandumOfUnderstanding(c *fiber.Ctx) error {
	var memorandumOfUnderstanding MemorandumOfUnderstanding
	if err := c.BodyParser(&memorandumOfUnderstanding); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed memorandum of understanding payload: %+v", memorandumOfUnderstanding)

	if err := h.Service.CreateMemorandumOfUnderstanding(&memorandumOfUnderstanding); err != nil {
		log.Printf("Error creating school: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create memorandum of understanding",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Memorandum of understanding created successfully",
		memorandumOfUnderstanding,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetMemorandumOfUnderstanding(c *fiber.Ctx) error {
	memorandumOfUnderstanding, err := h.Service.GetMemorandumOfUnderstanding()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get memorandum of understanding",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Memorandum of understanding retrieved successfully",
		memorandumOfUnderstanding,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetMemorandumOfUnderstandingByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid memorandum of understanding ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	memorandumOfUnderstanding, err := h.Service.GetMemorandumOfUnderstandingByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get memorandum of understanding",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Memorandum of understanding retrieved successfully",
		memorandumOfUnderstanding,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateMemorandumOfUnderstanding(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid memorandum of understanding ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var memorandumOfUnderstanding MemorandumOfUnderstanding
	if err := c.BodyParser(&memorandumOfUnderstanding); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	memorandumOfUnderstanding.ID = uint64(id)

	if err := h.Service.UpdateMemorandumOfUnderstanding(&memorandumOfUnderstanding); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Memorandum of understanding updated successfully",
		memorandumOfUnderstanding,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteMemorandumOfUnderstanding(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting memorandum of understanding ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid memorandum of understanding ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteMemorandumOfUnderstanding(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting memorandum of understanding by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Memorandum of understanding deleted successfully", nil, nil, nil, nil, nil,
	))
}
