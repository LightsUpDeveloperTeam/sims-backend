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

//Collaboration History Repositories

func (r *Repository) CreateCollaborationHistory(collaborationHistory *CollaborationHistory) error {
	return r.DB.Create(collaborationHistory).Error
}

func (r *Repository) GetCollaborationHistory() ([]CollaborationHistory, error) {
	var collaborationHistory []CollaborationHistory
	err := r.DB.Find(&collaborationHistory).Error
	return collaborationHistory, err
}

func (r *Repository) GetCollaborationHistoryByID(id uint64) (*CollaborationHistory, error) {
	var collaborationHistory CollaborationHistory
	err := r.DB.First(&collaborationHistory, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("collaboration history not found")
	}
	return &collaborationHistory, err
}

func (r *Repository) UpdateCollaborationHistory(collaborationHistory *CollaborationHistory) error {
	return r.DB.Save(collaborationHistory).Error
}

func (r *Repository) DeleteCollaborationHistory(id uint64, deletedBy uint64) error {
	log.Printf("Deleting collaboration history with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&CollaborationHistory{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}

//Memorandum Of Understanding Repositories

func (r *Repository) CreateMemorandumOfUnderstanding(memorandumOfUnderstanding *MemorandumOfUnderstanding) error {
	return r.DB.Create(memorandumOfUnderstanding).Error
}

func (r *Repository) GetAllMemorandumOfUnderstanding() ([]MemorandumOfUnderstanding, error) {
	var memorandumOfUnderstanding []MemorandumOfUnderstanding
	err := r.DB.Find(&memorandumOfUnderstanding).Error
	return memorandumOfUnderstanding, err
}

func (r *Repository) GetMemorandumOfUnderstandingByID(id uint64) (*MemorandumOfUnderstanding, error) {
	var memorandumOfUnderstanding MemorandumOfUnderstanding
	err := r.DB.First(&memorandumOfUnderstanding, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Memorandum of understanding not found")
	}
	return &memorandumOfUnderstanding, err
}

func (r *Repository) UpdateMemorandumOfUnderstanding(memorandumOfUnderstanding *MemorandumOfUnderstanding) error {
	return r.DB.Save(memorandumOfUnderstanding).Error
}

func (r *Repository) DeleteMemorandumOfUnderstanding(id uint64, deletedBy uint64) error {
	log.Printf("Deleting memorandum of understanding with ID: %d, deletedBy: %d", id, deletedBy)
	return r.DB.Model(&MemorandumOfUnderstanding{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deletedBy,
		"deleted_at": time.Now(),
	}).Error
}
