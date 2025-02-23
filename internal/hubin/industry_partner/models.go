package industry_partner

import "time"

type IndustryPartner struct {
	Id              uint      `gorm:"primary_key" json:"id"`
	PartnerName     string    `gorm:"not null" json:"partner_name"`
	WorkSector      string    `gorm:"not null" json:"work_sector"`
	Address         string    `gorm:"not null" json:"address"`
	Email           string    `gorm:"not null" json:"email"`
	TelephoneNumber string    `gorm:"not null" json:"telephone_number"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
