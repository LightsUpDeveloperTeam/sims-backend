package internship_information

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

//Internship Contract Repositories

func (r *Repository) CreateInternshipContract(internshipContract *InternshipContract) error {
	return r.DB.Create(internshipContract).Error
}

func (r *Repository) GetAllInternshipContract() ([]InternshipContract, error) {
	var internshipContract []InternshipContract
	err := r.DB.Find(&internshipContract).Error
	return internshipContract, err
}

func (r *Repository) GetInternshipContractByID(id uint64) (*InternshipContract, error) {
	var internshipContract InternshipContract
	err := r.DB.First(&internshipContract, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("internship contract not found")
	}
	return &internshipContract, err
}

func (r *Repository) UpdateInternshipContract(internshipContract *InternshipContract) error {
	return r.DB.Save(internshipContract).Error
}

func (r *Repository) DeleteInternshipContract(id uint64, deletedBy uint64) error {
	log.Printf("Deleting internship contract with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&InternshipContract{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}

//Internship Document Repositories

func (r *Repository) CreateInternshipDocument(internshipDocument *InternshipDocument) error {
	return r.DB.Create(internshipDocument).Error
}

func (r *Repository) GetAllInternshipDocument() ([]InternshipDocument, error) {
	var internshipDocument []InternshipDocument
	err := r.DB.Find(&internshipDocument).Error
	return internshipDocument, err
}

func (r *Repository) GetInternshipDocumentByID(id uint64) (*InternshipDocument, error) {
	var internshipDocument InternshipDocument
	err := r.DB.First(&internshipDocument, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("internship document not found")
	}
	return &internshipDocument, err
}

func (r *Repository) UpdateInternshipDocument(internshipDocument *InternshipDocument) error {
	return r.DB.Save(internshipDocument).Error
}

func (r *Repository) DeleteInternshipDocument(id uint64, deletedBy uint64) error {
	log.Printf("Deleting internship document with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&InternshipDocument{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}

//Internship Evaluation Repositories

func (r *Repository) CreateInternshipEvaluation(internshipEvaluation *InternshipEvaluation) error {
	return r.DB.Create(internshipEvaluation).Error
}

func (r *Repository) GetAllInternshipEvaluation() ([]InternshipEvaluation, error) {
	var internshipEvaluation []InternshipEvaluation
	err := r.DB.Find(&internshipEvaluation).Error
	return internshipEvaluation, err
}

func (r *Repository) GetInternshipEvaluationByID(id uint64) (*InternshipEvaluation, error) {
	var internshipEvaluation InternshipEvaluation
	err := r.DB.First(&internshipEvaluation, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("internship evaluation not found")
	}
	return &internshipEvaluation, err
}

func (r *Repository) UpdateInternshipEvaluation(internshipEvaluation *InternshipEvaluation) error {
	return r.DB.Save(internshipEvaluation).Error
}

func (r *Repository) DeleteInternshipEvaluation(id uint64, deletedBy uint64) error {
	log.Printf("Deleting internship contract with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&InternshipEvaluation{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}

//Internship Progress Repositories

func (r *Repository) CreateInternshipProgress(internshipProgress *InternshipProgress) error {
	return r.DB.Create(internshipProgress).Error
}

func (r *Repository) GetAllInternshipProgress() ([]InternshipProgress, error) {
	var internshipProgress []InternshipProgress
	err := r.DB.Find(&internshipProgress).Error
	return internshipProgress, err
}

func (r *Repository) GetInternshipProgressByID(id uint64) (*InternshipProgress, error) {
	var internshipProgress InternshipProgress
	err := r.DB.First(&internshipProgress, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("internship progress not found")
	}
	return &internshipProgress, err
}

func (r *Repository) UpdateInternshipProgress(internshipProgress *InternshipProgress) error {
	return r.DB.Save(internshipProgress).Error
}

func (r *Repository) DeleteInternshipProgress(id uint64, deletedBy uint64) error {
	log.Printf("Deleting internship progress with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&InternshipProgress{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}
