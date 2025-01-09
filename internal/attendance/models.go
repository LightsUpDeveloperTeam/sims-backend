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

type Shift struct {
	ID          uint   // ID shift
	Name        string // Nama shift
	Description string // Deskripsi shift
}

type ShiftSchedule struct {
	ID        uint      // ID shift schedule
	ShiftID   uint      // ID shift
	DayOfWeek int       // Hari kerja (0 = Minggu, 6 = Sabtu)
	StartTime time.Time // Waktu mulai shift
	EndTime   time.Time // Waktu selesai shift
}

type UserShiftAssignment struct {
	ID        uint      // ID assignment
	UserID    uint      // ID user
	ShiftID   uint      // ID shift
	StartDate time.Time // Tanggal mulai assignment
	EndDate   time.Time // Tanggal selesai assignment
}

type UserSchedule struct {
	ID        uint      // ID schedule
	UserID    uint      // ID user
	DayOfWeek int       // Hari kerja
	StartTime time.Time // Waktu mulai kerja
	EndTime   time.Time // Waktu selesai kerja
}

type AttendanceRecord struct {
	ID                  uint       `gorm:"primaryKey"`
	UserID              uint       `gorm:"not null"`
	Date                time.Time  `gorm:"type:date;not null"`
	ClockInTime         *time.Time `gorm:"type:datetime"`
	ClockOutTime        *time.Time `gorm:"type:datetime"`
	ClockInLatitude     float64    `gorm:"type:decimal(9,6)"`
	ClockInLongitude    float64    `gorm:"type:decimal(9,6)"`
	ClockOutLatitude    float64    `gorm:"type:decimal(9,6)"`
	ClockOutLongitude   float64    `gorm:"type:decimal(9,6)"`
	EarlyClockOutReason string     // Alasan clock-out lebih awal
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Holiday struct {
	ID                   uint      // ID holiday
	Name                 string    // Nama hari libur
	Type                 string    // Tipe hari libur (religious, national, international)
	Date                 time.Time // Tanggal hari libur
	IsHoliday            bool      // Apakah hari itu libur
	IsCelebratedAtSchool bool      // Apakah dirayakan di sekolah
}
