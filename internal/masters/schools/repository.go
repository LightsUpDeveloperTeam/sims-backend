package schoolsmasterdata

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateSchool(school *School) error {
	return r.DB.Create(school).Error
}

func (r *Repository) GetAllSchools() ([]School, error) {
	var schools []School
	err := r.DB.Find(&schools).Error
	return schools, err
}

func (r *Repository) GetSchoolByID(id uint64) (*School, error) {
	var school School
	err := r.DB.First(&school, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("school not found")
	}
	return &school, err
}

func (r *Repository) UpdateSchool(school *School) error {
	return r.DB.Save(school).Error
}

func (r *Repository) DeleteSchool(id uint64) error {
	return r.DB.Delete(&School{ID: id}).Error
}