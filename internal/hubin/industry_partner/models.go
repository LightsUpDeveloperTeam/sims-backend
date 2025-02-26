package industry_partner

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type IndustryPartner struct {
	ID              uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	PartnerName     string          `gorm:"size:255;not null" json:"partner_name"`
	FieldOfWork     json.RawMessage `gorm:"type:jsonb" json:"field_of_work"`
	Address         string          `gorm:"size:255;not null;unique" json:"address"`
	Email           string          `gorm:"size:50" json:"email"`
	TelephoneNumber string          `gorm:"size:255" json:"telephone_number"`
	DeletedBy       *uint64         `json:"deleted_by"`
	DeletedAt       *time.Time      `json:"deleted_at"`
	CreatedAt       time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}

func (IndustryPartner) TableName() string {
	return "industry_partner"
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
