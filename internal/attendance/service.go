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

func (s *AttendanceService) ClockInOut(userID uint, latitude, longitude float64, earlyReason string) (*AttendanceRecord, error) {
	// Get today's date
	date := time.Now().Format("2006-01-02")
	record, err := s.Repo.GetAttendanceByUserAndDate(userID, date)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	if record == nil {
		// Clock-In
		shift, err := s.Repo.GetShiftScheduleByUser(userID, int(now.Weekday()))
		if err != nil {
			return nil, err
		}

		if shift == nil {
			return nil, errors.New("no shift schedule found for the user today")
		}

		if now.Before(shift.StartTime) {
			return nil, errors.New("cannot clock in before shift start time")
		}

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

	if record.ClockOutTime != nil {
		return nil, errors.New("user has already clocked out today")
	}

	// Clock-Out
	shift, err := s.Repo.GetShiftScheduleByUser(userID, int(now.Weekday()))
	if err != nil {
		return nil, err
	}

	if shift == nil {
		return nil, errors.New("no shift schedule found for the user today")
	}

	if now.Before(shift.EndTime) && earlyReason == "" {
		return nil, errors.New("clocking out earlier than scheduled, reason required")
	}

	record.ClockOutTime = &now
	record.ClockOutLatitude = latitude
	record.ClockOutLongitude = longitude

	if now.Before(shift.EndTime) {
		record.EarlyClockOutReason = earlyReason
	}

	err = s.Repo.UpdateAttendance(record)
	if err != nil {
		return nil, err
	}

	return record, nil
}
