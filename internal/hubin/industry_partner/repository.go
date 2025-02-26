package industry_partner

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateIndustryPartner(industryPartner *IndustryPartner) error {
	return r.DB.Create(industryPartner).Error
}

func (r *Repository) GetAllIndustryPartner() ([]IndustryPartner, error) {
	var industryPartner []IndustryPartner
	err := r.DB.Find(&industryPartner).Error
	return industryPartner, err
}

func (r *Repository) GetIndustryPartnerByID(id uint64) (*IndustryPartner, error) {
	var industryPartner IndustryPartner
	err := r.DB.First(&industryPartner, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("industry partner not found")
	}
	return &industryPartner, err
}

func (r *Repository) UpdateIndustryPartner(industryPartner *IndustryPartner) error {
	return r.DB.Save(industryPartner).Error
}

func (r *Repository) DeleteIndustryPartner(id uint64, deletedBy uint64) error {
	log.Printf("Deleting user with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&IndustryPartner{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}
