package attendance

import (
	"errors"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	DB *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{DB: db}
}

func (r *AttendanceRepository) CreateAttendance(record *AttendanceRecord) error {
	return r.DB.Create(record).Error
}

func (r *AttendanceRepository) UpdateAttendance(record *AttendanceRecord) error {
	return r.DB.Save(record).Error
}

func (r *AttendanceRepository) GetAttendanceByUserAndDate(userID uint, date string) (*AttendanceRecord, error) {
	var record AttendanceRecord
	err := r.DB.Where("user_id = ? AND date = ?", userID, date).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &record, err
}
