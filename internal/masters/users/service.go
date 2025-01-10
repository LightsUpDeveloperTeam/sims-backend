package usersmasterdata

import (
	"errors"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	if user.Email == "" {
		return errors.New("email are required")
	}
	return s.Repo.CreateUser(user)
}

func (s *Service) GetUserByID(id uint64) (*User, error) {
	return s.Repo.GetUserByID(id)
}

func (s *Service) UpdateUser(user *User) error {
	return s.Repo.UpdateUser(user)
}

func (s *Service) DeleteUser(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteUser(id, deletedBy)
}

func (s *Service) GetAllUsers(filter map[string]interface{}) ([]User, error) {
	return s.Repo.GetAllUsers(filter)
}

func (s *Service) CreateRole(role *Role) error {
	if role.Name == "" {
		return errors.New("role name is required")
	}
	return s.Repo.CreateRole(role)
}

func (s *Service) UpdateRole(role *Role) error {
	return s.Repo.UpdateRole(role)
}

func (s *Service) DeleteRole(id uint64) error {
	return s.Repo.DeleteRole(id)
}

func (s *Service) CreatePermission(permission *Permission) error {
	if permission.Name == "" {
		return errors.New("permission name is required")
	}
	return s.Repo.CreatePermission(permission)
}

func (s *Service) AssignPermissionToRole(roleID, permissionID uint64) error {
	return s.Repo.AssignPermissionToRole(roleID, permissionID)
}

func (s *Service) GetPermissionsByRole(roleID uint64) ([]Permission, error) {
	return s.Repo.GetPermissionsByRole(roleID)
}
