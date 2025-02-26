package internship_information

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

//Internship Contract Handler

func (h *Handler) CreateInternshipContract(c *fiber.Ctx) error {
	var internshipContract InternshipContract
	if err := c.BodyParser(&internshipContract); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed internship contract payload: %+v", internshipContract)

	if err := h.Service.CreateInternshipContract(&internshipContract); err != nil {
		log.Printf("Error creating internship contract: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create internship contract",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship contract created successfully",
		internshipContract,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipContract(c *fiber.Ctx) error {
	internshipContract, err := h.Service.GetInternshipContract()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship contract",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship contract retrieved successfully",
		internshipContract,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipContractByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship contract ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	internshipContract, err := h.Service.GetInternshipContractByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship contract",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship contract retrieved successfully",
		internshipContract,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateInternshipContract(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship contract ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var internshipContract InternshipContract
	if err := c.BodyParser(&internshipContract); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	internshipContract.ID = uint64(id)

	if err := h.Service.UpdateInternshipContract(&internshipContract); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship contract updated successfully",
		internshipContract,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteInternshipContract(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting internship contract ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid internship contract ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteInternshipContract(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting internship contract by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Internship contract deleted successfully", nil, nil, nil, nil, nil,
	))
}

//Internship Document Handler

func (h *Handler) CreateInternshipDocument(c *fiber.Ctx) error {
	var internshipDocument InternshipDocument
	if err := c.BodyParser(&internshipDocument); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed internship document payload: %+v", internshipDocument)

	if err := h.Service.CreateInternshipDocument(&internshipDocument); err != nil {
		log.Printf("Error creating internship document: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create internship document",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship document created successfully",
		internshipDocument,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipDocument(c *fiber.Ctx) error {
	internshipDocument, err := h.Service.GetInternshipDocument()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship document",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship document retrieved successfully",
		internshipDocument,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipDocumentByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship document ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	internshipDocument, err := h.Service.GetInternshipDocumentByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship document",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship document retrieved successfully",
		internshipDocument,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateInternshipDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship document ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var internshipDocument InternshipDocument
	if err := c.BodyParser(&internshipDocument); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	internshipDocument.ID = uint64(id)

	if err := h.Service.UpdateInternshipDocument(&internshipDocument); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship document updated successfully",
		internshipDocument,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteInternshipDocument(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting internship document ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid internship document ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteInternshipDocument(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting internship document by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Internship document deleted successfully", nil, nil, nil, nil, nil,
	))
}

//Internship Evaluation Handler

func (h *Handler) CreateInternshipEvaluation(c *fiber.Ctx) error {
	var internshipEvaluation InternshipEvaluation
	if err := c.BodyParser(&internshipEvaluation); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed internship evaluation payload: %+v", internshipEvaluation)

	if err := h.Service.CreateInternshipEvaluation(&internshipEvaluation); err != nil {
		log.Printf("Error creating internship evaluation: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create internship evaluation",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship contract evaluation successfully",
		internshipEvaluation,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipEvaluation(c *fiber.Ctx) error {
	internshipEvaluation, err := h.Service.GetInternshipEvaluation()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship evaluation",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship evaluation retrieved successfully",
		internshipEvaluation,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipEvaluationByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship evaluation ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	internshipEvaluation, err := h.Service.GetInternshipEvaluationByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship evaluation",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship evaluation retrieved successfully",
		internshipEvaluation,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateInternshipEvaluation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship evaluation ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var internshipEvaluation InternshipEvaluation
	if err := c.BodyParser(&internshipEvaluation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	internshipEvaluation.ID = uint64(id)

	if err := h.Service.UpdateInternshipEvaluation(&internshipEvaluation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship evaluation updated successfully",
		internshipEvaluation,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteInternshipEvaluation(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting internship evaluation ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid internship evaluation ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteInternshipEvaluation(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting internship evaluation by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Internship evaluation deleted successfully", nil, nil, nil, nil, nil,
	))
}

//Internship Progress Handler

func (h *Handler) CreateInternshipProgress(c *fiber.Ctx) error {
	var internshipProgress InternshipProgress
	if err := c.BodyParser(&internshipProgress); err != nil {
		log.Printf("Body parsing error: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	log.Printf("Parsed internship progress payload: %+v", internshipProgress)

	if err := h.Service.CreateInternshipProgress(&internshipProgress); err != nil {
		log.Printf("Error creating internship progress: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to create internship progress",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship progress created successfully",
		internshipProgress,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipProgress(c *fiber.Ctx) error {
	internshipProgress, err := h.Service.GetInternshipProgress()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship progress",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship progress retrieved successfully",
		internshipProgress,
		nil, nil, nil, nil,
	))
}

func (h *Handler) GetInternshipProgressByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship progress ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	internshipProgress, err := h.Service.GetInternshipProgressByID(uint64(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to get internship progress",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship progress retrieved successfully",
		internshipProgress,
		nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateInternshipProgress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid internship progress ID",
			nil,
			nil, nil, nil, nil,
		))
	}

	var internshipProgress InternshipProgress
	if err := c.BodyParser(&internshipProgress); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}
	internshipProgress.ID = uint64(id)

	if err := h.Service.UpdateInternshipProgress(&internshipProgress); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS",
		"Internship progress updated successfully",
		internshipProgress,
		nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteInternshipProgress(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting internship progress ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid internship progress ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteInternshipProgress(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting internship progress by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Internship progress deleted successfully", nil, nil, nil, nil, nil,
	))
}
