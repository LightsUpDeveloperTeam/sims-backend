package usersmasterdata

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) GetUserByID(id uint64) (*User, error) {
	var user User
	err := r.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *Repository) UpdateUser(user *User) error {
	return r.DB.Save(user).Error
}

func (r *Repository) DeleteUser(id uint64, deletedBy uint64) error {
	log.Printf("Deleting user with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}

func (r *Repository) GetAllUsers(filter map[string]interface{}) ([]User, error) {
	var users []User
	query := r.DB.Preload("Roles").Preload("School")
	for key, value := range filter {
		query = query.Where(key+" = ?", value)
	}
	err := query.Find(&users).Error
	return users, err
}

func (r *Repository) CreateRole(role *Role) error {
	return r.DB.Create(role).Error
}

func (r *Repository) UpdateRole(role *Role) error {
	return r.DB.Save(role).Error
}

func (r *Repository) DeleteRole(id uint64) error {
	return r.DB.Delete(&Role{}, id).Error
}

func (r *Repository) CreatePermission(permission *Permission) error {
	return r.DB.Create(permission).Error
}

func (r *Repository) AssignPermissionToRole(roleID, permissionID uint64) error {
	return r.DB.Create(&RolePermission{
		RoleID:       roleID,
		PermissionID: permissionID,
		CreatedAt:    time.Now(),
	}).Error
}

func (r *Repository) GetPermissionsByRole(roleID uint64) ([]Permission, error) {
	var permissions []Permission
	err := r.DB.Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Where("role_permissions.role_id = ?", roleID).Find(&permissions).Error
	return permissions, err
}
