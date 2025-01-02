package schoolsmasterdata

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type School struct {
	ID                   uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                 string    `gorm:"size:255;not null" json:"name"`
	Address              string    `gorm:"type:text" json:"address"`
	ContactEmail         string    `gorm:"size:255;not null;unique" json:"contact_email"`
	ContactPhone         string    `gorm:"size:50" json:"contact_phone"`
	LogoURL              string    `gorm:"size:255" json:"logo_url"`
	SubscriptionStatus   string    `gorm:"size:20;not null;default:'trial'" json:"subscription_status"` 
	SubscriptionPlan     string    `gorm:"size:255" json:"subscription_plan"`
	SubscriptionExpiryDate *CustomDate `json:"subscription_expiry_date"`
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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