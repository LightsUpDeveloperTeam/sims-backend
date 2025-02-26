package BKK

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

// Internship Vacancy Services

func (s *Service) CreateVacancy(vacancy *InternshipVacancy) error {
	if vacancy.PositionName == "" || vacancy.Description == "" {
		return errors.New("Fill all the vacancy details!")
	}

	return s.Repo.CreateVacancy(vacancy)
}

func (s *Service) GetVacancies() ([]InternshipVacancy, error) {
	return s.Repo.GetAllVacancy()
}

func (s *Service) GetVacancyByID(id uint64) (*InternshipVacancy, error) {
	return s.Repo.GetVacancyByID(id)
}

func (s *Service) UpdateVacancy(vacancy *InternshipVacancy) error {
	return s.Repo.UpdateVacancy(vacancy)
}

func (s *Service) DeleteVacancy(id uint64, deletedBy uint64) error {
	fmt.Println("Service received ID:", id)
	return s.Repo.DeleteVacancy(id, deletedBy)
}

// Registration Vacancy Services

func (s *Service) CreateRegistration(registration *InternshipRegistration) error {
	if registration.CompanyName == "" || registration.Position == "" || registration.RegistrationDate == nil {
		return errors.New("Fill all the registration details!")
	}

	return s.Repo.CreateRegistration(registration)
}

func (s *Service) GetRegistrations() ([]InternshipRegistration, error) {
	return s.Repo.GetAllRegistration()
}

func (s *Service) GetRegistrationByID(id uint64) (*InternshipRegistration, error) {
	return s.Repo.GetRegistrationByID(id)
}

func (s *Service) DeleteRegistration(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteRegistration(id, deletedBy)
}

// Alumnus Distribution Services

func (s *Service) CreateDistribution(distribution *AlumnusDistribution) error {
	if distribution.CompanyName == "" || distribution.Position == "" {
		return errors.New("Fill all the distribution details!")
	}

	return s.Repo.CreateDistribution(distribution)
}

func (s *Service) GetDistribution() ([]AlumnusDistribution, error) {
	return s.Repo.GetAllDistribution()
}

func (s *Service) GetDistributionByID(id uint64) (*AlumnusDistribution, error) {
	return s.Repo.GetDistributionByID(id)
}

func (s *Service) UpdateDistribution(alumnusDistribution *AlumnusDistribution) error {
	return s.Repo.UpdateDistribution(alumnusDistribution)
}

func (s *Service) DeleteDistribution(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteDistribution(id, deletedBy)
}
