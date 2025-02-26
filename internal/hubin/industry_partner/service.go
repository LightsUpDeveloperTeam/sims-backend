package industry_partner

import (
	"errors"
	"fmt"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateIndustryPartner(industryPartner *IndustryPartner) error {
	if industryPartner.PartnerName == "" || industryPartner.Email == "" || industryPartner.FieldOfWork == nil {
		return errors.New("Fill all the creation details!")
	}

	return s.Repo.CreateIndustryPartner(industryPartner)
}

func (s *Service) GetIndustryPartner() ([]IndustryPartner, error) {
	return s.Repo.GetAllIndustryPartner()
}

func (s *Service) GetIndustryPartnerByID(id uint64) (*IndustryPartner, error) {
	return s.Repo.GetIndustryPartnerByID(id)
}

func (s *Service) UpdateIndustryPartner(industryPartner *IndustryPartner) error {
	return s.Repo.UpdateIndustryPartner(industryPartner)
}

func (s *Service) DeleteIndustryPartner(id uint64, deletedBy uint64) error {
	fmt.Println("DeleteIndustryPartner id:", id)
	return s.Repo.DeleteIndustryPartner(id, deletedBy)
}
