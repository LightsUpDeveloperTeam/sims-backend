package BKK

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type InternshipVacancy struct {
	Id                uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	IndustryPartnerId uint64      `gorm:"primaryKey" json:"industry_partner_id"`
	PositionName      string      `gorm:"not null" json:"position_name"`
	Description       string      `gorm:"not null" json:"description"`
	OpenDate          *CustomDate `gorm:"not null" json:"open_date"`
	CreatedAt         time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type InternshipRegistration struct {
	Id                  uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentId           uint64      `gorm:"primaryKey" json:"student_id"`
	InternshipVacancyId uint64      `gorm:"primaryKey" json:"internship_vacancy_id"`
	CompanyName         string      `gorm:"not null" json:"company_name"`
	Position            string      `gorm:"not null" json:"position"`
	RegistrationDate    *CustomDate `gorm:"not null" json:"registration_date"`
	CreatedAt           time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type AlumnusDistribution struct {
	Id          uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentId   uint64      `gorm:"primaryKey" json:"student_id"`
	CompanyName string      `gorm:"not null" json:"company_name"`
	Position    string      `gorm:"not null" json:"position"`
	StartDate   *CustomDate `gorm:"not null" json:"start_date"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type CustomDate struct {
	time.Time
}

// MarshalJSON memformat tanggal sebagai "YYYY-MM-DD"
func (c CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.Format("2006-01-02"))
}

// UnmarshalJSON memparsing tanggal dalam format "YYYY-MM-DD"
func (c *CustomDate) UnmarshalJSON(data []byte) error {
	strInput := string(data)
	parsedTime, err := time.Parse(`"2006-01-02"`, strInput)
	if err != nil {
		return errors.New("invalid date format, use YYYY-MM-DD")
	}
	c.Time = parsedTime
	return nil
}

// Value mengimplementasikan sql.Valuer untuk CustomDate
func (c CustomDate) Value() (driver.Value, error) {
	return c.Time, nil
}

func (c *CustomDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		c.Time = v
		return nil
	default:
		return errors.New("invalid type for CustomDate")
	}
}
