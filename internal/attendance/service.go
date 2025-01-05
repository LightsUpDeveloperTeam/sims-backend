package attendance

import (
	"errors"
	"time"
)

type AttendanceService struct {
	Repo *AttendanceRepository
}

func NewAttendanceService(repo *AttendanceRepository) *AttendanceService {
	return &AttendanceService{Repo: repo}
}

func (s *AttendanceService) ClockIn(userID uint, latitude, longitude float64) (*AttendanceRecord, error) {
	// Periksa apakah user sudah clock-in hari ini
	date := time.Now().Format("2006-01-02")
	record, err := s.Repo.GetAttendanceByUserAndDate(userID, date)
	if err != nil {
		return nil, err
	}
	if record != nil {
		return nil, errors.New("user has already clocked in today")
	}

	// Buat catatan clock-in
	now := time.Now()
	newRecord := &AttendanceRecord{
		UserID:           userID,
		Date:             now,
		ClockInTime:      &now,
		ClockInLatitude:  latitude,
		ClockInLongitude: longitude,
	}
	err = s.Repo.CreateAttendance(newRecord)
	if err != nil {
		return nil, err
	}

	return newRecord, nil
}

func (s *AttendanceService) ClockOut(userID uint, latitude, longitude float64) (*AttendanceRecord, error) {
	// Ambil catatan absensi untuk hari ini
	date := time.Now().Format("2006-01-02")
	record, err := s.Repo.GetAttendanceByUserAndDate(userID, date)
	if err != nil {
		return nil, err
	}
	if record == nil {
		return nil, errors.New("no clock-in record found for today")
	}

	// Perbarui catatan clock-out
	now := time.Now()
	record.ClockOutTime = &now
	record.ClockOutLatitude = latitude
	record.ClockOutLongitude = longitude

	err = s.Repo.UpdateAttendance(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}
