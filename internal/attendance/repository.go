package attendance

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// AttendanceRepository handles database operations for attendance.
type AttendanceRepository struct {
	DB *gorm.DB
}

// NewAttendanceRepository creates a new instance of AttendanceRepository.
func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{DB: db}
}

// GetShiftScheduleByUser fetches the shift schedule for a specific user and day of the week.
// Parameters:
// - userID: ID of the user
// - dayOfWeek: Day of the week (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
// Returns:
// - *ShiftSchedule: The shift schedule if found
// - error: Any error that occurred during the operation
func (r *AttendanceRepository) GetShiftScheduleByUser(userID uint, dayOfWeek int) (*ShiftSchedule, error) {
	if userID == 0 || dayOfWeek < 0 || dayOfWeek > 6 {
		return nil, errors.New("invalid userID or dayOfWeek")
	}

	var schedule ShiftSchedule
	err := r.DB.Joins("JOIN user_shift_assignments ON user_shift_assignments.shift_id = shift_schedules.shift_id").
		Where("user_shift_assignments.user_id = ? AND shift_schedules.day_of_week = ?", userID, dayOfWeek).
		First(&schedule).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("No shift schedule found for userID %d on day %d", userID, dayOfWeek)
		return nil, nil
	}
	if err != nil {
		log.Printf("Error fetching shift schedule for user %d on day %d: %v", userID, dayOfWeek, err)
	}
	return &schedule, err
}

// CreateAttendance creates a new attendance record in the database.
func (r *AttendanceRepository) CreateAttendance(record *AttendanceRecord) error {
	if record == nil {
		return errors.New("attendance record cannot be nil")
	}
	err := r.DB.Create(record).Error
	if err != nil {
		log.Printf("Error creating attendance record: %v", err)
	}
	return err
}

// CheckUserExists checks if a user exists in the database.
func (r *AttendanceRepository) CheckUserExists(userID uint) (bool, error) {
	if userID == 0 {
		return false, errors.New("invalid userID")
	}

	var user User
	err := r.DB.First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil // User does not exist
	}
	if err != nil {
		log.Printf("Error checking user existence for userID %d: %v", userID, err)
	}
	return err == nil, err
}

// UpdateAttendance updates an existing attendance record in the database.
func (r *AttendanceRepository) UpdateAttendance(record *AttendanceRecord) error {
	if record == nil || record.ID == 0 {
		return errors.New("invalid attendance record")
	}

	err := r.DB.Save(record).Error
	if err != nil {
		log.Printf("Error updating attendance record with ID %d: %v", record.ID, err)
	}
	return err
}

// GetAttendanceByUserAndDate fetches an attendance record for a specific user on a given date.
// Parameters:
// - userID: ID of the user
// - date: Date in the format "YYYY-MM-DD"
// Returns:
// - *AttendanceRecord: The attendance record if found
// - error: Any error that occurred during the operation
func (r *AttendanceRepository) GetAttendanceByUserAndDate(userID uint, date string) (*AttendanceRecord, error) {
	if userID == 0 || date == "" {
		return nil, errors.New("invalid userID or date")
	}

	var record AttendanceRecord
	err := r.DB.Where("user_id = ? AND date = ?", userID, date).First(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // No attendance record found
	}
	if err != nil {
		log.Printf("Error fetching attendance record for user %d on date %s: %v", userID, date, err)
	}
	return &record, err
}
