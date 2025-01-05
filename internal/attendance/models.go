package attendance

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"unique"`
	Role      string `gorm:"type:enum('student', 'teacher', 'school_admissions')"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AttendanceRecord struct {
	ID                uint       `gorm:"primaryKey"`
	UserID            uint       `gorm:"not null"`
	Date              time.Time  `gorm:"type:date;not null"`
	ClockInTime       *time.Time `gorm:"type:datetime"`
	ClockOutTime      *time.Time `gorm:"type:datetime"`
	ClockInLatitude   float64    `gorm:"type:decimal(9,6)"`
	ClockInLongitude  float64    `gorm:"type:decimal(9,6)"`
	ClockOutLatitude  float64    `gorm:"type:decimal(9,6)"`
	ClockOutLongitude float64    `gorm:"type:decimal(9,6)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
