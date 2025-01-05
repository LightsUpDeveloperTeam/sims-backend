package usersmasterdata

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

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user User
	log.Printf("Received request to create user: %+v", user)
	if err := c.BodyParser(&user); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	if err := h.Service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "User created successfully", user, nil, nil, nil, nil,
	))
}

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting user ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Received request to get user by ID: %d", id)
	user, err := h.Service.GetUserByID(uint64(id))
	if err != nil {
		log.Printf("Error retrieving user by ID: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Retrieved user by ID: %+v", user)
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "User retrieved successfully", user, nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting user ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Received request to update user: %d", id)
	var user User
	if err := c.BodyParser(&user); err != nil {
		log.Printf("Error parsing request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	user.ID = uint64(id)
	if err := h.Service.UpdateUser(&user); err != nil {
		log.Printf("Error updating user: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	log.Printf("Updated user: %+v", user)
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "User updated successfully", user, nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Error getting user ID from request: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid user ID", nil, nil, nil, nil, nil,
		))
	}

	deletedBy := uint64(1)

	err = h.Service.DeleteUser(uint64(id), deletedBy)
	if err != nil {
		log.Printf("Error deleting user by ID: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,))
	}

	return c.JSON(utils.CreateResponse(
		"SUCCESS", "User deleted successfully", nil, nil, nil, nil, nil,
	))
}

func (h *Handler) CreateRole(c *fiber.Ctx) error {
	var role Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	if err := h.Service.CreateRole(&role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "Role created successfully", role, nil, nil, nil, nil,
	))
}

func (h *Handler) UpdateRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid role ID", nil, nil, nil, nil, nil,
		))
	}
	var role Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	role.ID = uint64(id)
	if err := h.Service.UpdateRole(&role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Role updated successfully", role, nil, nil, nil, nil,
	))
}

func (h *Handler) DeleteRole(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid role ID", nil, nil, nil, nil, nil,
		))
	}
	if err := h.Service.DeleteRole(uint64(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Role deleted successfully", nil, nil, nil, nil, nil,
	))
}

func (h *Handler) CreatePermission(c *fiber.Ctx) error {
	var permission Permission
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	if err := h.Service.CreatePermission(&permission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.Status(fiber.StatusCreated).JSON(utils.CreateResponse(
		"SUCCESS", "Permission created successfully", permission, nil, nil, nil, nil,
	))
}

func (h *Handler) AssignPermissionToRole(c *fiber.Ctx) error {
	var req struct {
		RoleID       uint64 `json:"role_id"`
		PermissionID uint64 `json:"permission_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}
	if err := h.Service.AssignPermissionToRole(req.RoleID, req.PermissionID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}
	return c.JSON(utils.CreateResponse(
		"SUCCESS", "Permission assigned to role successfully", nil, nil, nil, nil, nil,
	))
}

