package industry_partner

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type CollaborationHistory struct {
	ID                uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	IndustryPartnerID uint64      `gorm:"not null" json:"industry_partner_id"`
	StartDate         *CustomDate `json:"start_date"`
	EndDate           *CustomDate `json:"end_date"`
	Description       string      `gorm:"size:255;" json:"description"`
	Status            string      `gorm:"size:255;not null;default:'inactive'" json:"status"`
	DeletedBy         *uint64     `json:"deleted_by"`
	DeletedAt         *time.Time  `json:"deleted_at"`
	CreatedAt         time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type MemorandumOfUnderstanding struct {
	ID                uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	IndustryPartnerID uint64      `gorm:"not null" json:"industry_partner_id"`
	MOUNumber         string      `gorm:"size:255;" json:"mou_number"`
	StartDate         *CustomDate `json:"start_date"`
	EndDate           *CustomDate `json:"end_date"`
	Description       string      `gorm:"size:255;" json:"description"`
	MOUFile           string      `gorm:"size:255;not null;" json:"mou_file"`
	DeletedBy         *uint64     `json:"deleted_by"`
	DeletedAt         *time.Time  `json:"deleted_at"`
	CreatedAt         time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

func (CollaborationHistory) TableName() string {
	return "collaboration_history"
}

func (MemorandumOfUnderstanding) TableName() string {
	return "memorandum_of_understanding"
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
