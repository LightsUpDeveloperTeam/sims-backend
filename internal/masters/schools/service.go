package schoolsmasterdata

import "errors"

var validSubscriptionStatuses = map[string]bool{
	"active":   true,
	"inactive": true,
	"trial":    true,
	"expired":  true,
}

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateSchool(school *School) error {
	if school.Name == "" || school.ContactEmail == "" {
		return errors.New("name and contact email are required")
	}

	if !validSubscriptionStatuses[school.SubscriptionStatus] {
		return errors.New("invalid subscription status")
	}

	return s.Repo.CreateSchool(school)
}

func (s *Service) GetSchools() ([]School, error) {
	return s.Repo.GetAllSchools()
}

func (s *Service) GetSchoolByID(id uint64) (*School, error) {
	return s.Repo.GetSchoolByID(id)
}

func (s *Service) UpdateSchool(school *School) error {
	return s.Repo.UpdateSchool(school)
}

func (s *Service) DeleteSchool(id uint64) error {
	return s.Repo.DeleteSchool(id)
}