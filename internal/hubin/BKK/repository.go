package BKK

import (
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// Internship Vacancy Repositories

func (r *Repository) CreateVacancy(internshipVacancy *InternshipVacancy) error {
	return r.DB.Create(internshipVacancy).Error
}

func (r *Repository) GetAllVacancy() ([]InternshipVacancy, error) {
	var internshipVacancy []InternshipVacancy
	err := r.DB.Find(&internshipVacancy).Error
	return internshipVacancy, err
}

func (r *Repository) GetVacancyByID(id uint64) (*InternshipVacancy, error) {
	var idInternshipVacancy InternshipVacancy
	err := r.DB.First(&idInternshipVacancy, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Vacancy not found!")
	}
	return &idInternshipVacancy, err
}

func (r *Repository) UpdateVacancy(internshipVacancy *InternshipVacancy) error {
	return r.DB.Save(internshipVacancy).Error
}

func (r *Repository) DeleteVacancy(id uint64) error {
	return r.DB.Delete(&InternshipVacancy{Id: id}).Error
}

// Internship Registration Repositories

func (r *Repository) CreateRegistration(internshipRegistration *InternshipRegistration) error {
	return r.DB.Create(internshipRegistration).Error
}

func (r *Repository) GetAllRegistration() ([]InternshipRegistration, error) {
	var internshipRegistration []InternshipRegistration
	err := r.DB.Find(&internshipRegistration).Error
	return internshipRegistration, err
}

func (r *Repository) GetRegistrationByID(id uint64) (*InternshipRegistration, error) {
	var idInternshipRegistration InternshipRegistration
	err := r.DB.First(&idInternshipRegistration, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Registration not found!")
	}
	return &idInternshipRegistration, err
}

func (r *Repository) DeleteRegistration(id uint64) error {
	return r.DB.Delete(&InternshipRegistration{Id: id}).Error
}

// Alumnus Distribution Repositories

func (r *Repository) CreateDistribution(alumnusDistribution *AlumnusDistribution) error {
	return r.DB.Create(alumnusDistribution).Error
}

func (r *Repository) GetAllDistribution() ([]AlumnusDistribution, error) {
	var alumnusDistribution []AlumnusDistribution
	err := r.DB.Find(&alumnusDistribution).Error
	return alumnusDistribution, err
}

func (r *Repository) GetDistributionByID(id uint64) (*AlumnusDistribution, error) {
	var idAlumnusDistribution AlumnusDistribution
	err := r.DB.First(&idAlumnusDistribution, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Distribution not found!")
	}
	return &idAlumnusDistribution, err
}

func (r *Repository) UpdateDistribution(alumnusDistribution *AlumnusDistribution) error {
	return r.DB.Save(alumnusDistribution).Error
}

func (r *Repository) DeleteDistribution(id uint64) error {
	return r.DB.Delete(&AlumnusDistribution{Id: id}).Error
}
