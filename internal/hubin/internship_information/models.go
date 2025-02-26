package internship_information

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type InternshipContract struct {
	ID                  uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID           uint64      `gorm:"not null" json:"student_id"`
	CompanyName         string      `gorm:"size:255;not null" json:"company_name"`
	StartDate           *CustomDate `json:"start_date"`
	EndDate             *CustomDate `json:"end_date"`
	ContractDescription string      `gorm:"size:255;not null" json:"contract_description"`
	DeletedBy           *uint64     `json:"deleted_by"`
	DeletedAt           *time.Time  `json:"deleted_at"`
	CreatedAt           time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type InternshipDocument struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID    uint64     `gorm:"not null" json:"student_id"`
	Document     string     `gorm:"size:255;not null; default:'internship_report'" json:"document"`
	DocumentFile string     `gorm:"size:255;not null" json:"document_file"`
	Description  string     `gorm:"size:255;not null" json:"description"`
	DeletedBy    *uint64    `json:"deleted_by"`
	DeletedAt    *time.Time `json:"deleted_at"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

type InternshipEvaluation struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID uint64      `gorm:"not null" json:"student_id"`
	Date      *CustomDate `json:"date"`
	Rating    string      `gorm:"size:255;not null; default:'Bad'" json:"rating"`
	DeletedBy *uint64     `json:"deleted_by"`
	DeletedAt *time.Time  `json:"deleted_at"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

type InternshipProgress struct {
	ID                  uint64      `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentID           uint64      `gorm:"not null" json:"student_id"`
	Date                *CustomDate `json:"date"`
	DescriptionProgress string      `gorm:"size:255;not null" json:"description_progress"`
	Status              string      `gorm:"size:255;not null; default:'Bad'" json:"status"`
	DeletedBy           *uint64     `json:"deleted_by"`
	DeletedAt           *time.Time  `json:"deleted_at"`
	CreatedAt           time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}

func (InternshipContract) TableName() string {
	return "internship_contract"
}

func (InternshipDocument) TableName() string {
	return "internship_document"
}

func (InternshipEvaluation) TableName() string {
	return "internship_evaluation"
}

func (InternshipProgress) TableName() string {
	return "internship_progress"
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
