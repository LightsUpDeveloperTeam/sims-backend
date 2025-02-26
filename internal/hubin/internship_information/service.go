package internship_information

import "errors"

var documentType = map[string]bool{
	"commission":        true,
	"internship_report": true,
}

var evaluationRating = map[string]bool{
	"Good": true,
	"Bad":  true,
}

var progressStatus = map[string]bool{
	"Open":    true,
	"OnGoing": true,
	"Finish":  true,
}

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

//internship Contract Services

func (s *Service) CreateInternshipContract(internshipContract *InternshipContract) error {
	if internshipContract.CompanyName == "" || internshipContract.ContractDescription == "" {
		return errors.New("CompanyName or ContractDescription is required")
	}

	return s.Repo.CreateInternshipContract(internshipContract)
}

func (s *Service) GetInternshipContract() ([]InternshipContract, error) {
	return s.Repo.GetAllInternshipContract()
}

func (s *Service) GetInternshipContractByID(id uint64) (*InternshipContract, error) {
	return s.Repo.GetInternshipContractByID(id)
}

func (s *Service) UpdateInternshipContract(internshipContract *InternshipContract) error {
	return s.Repo.UpdateInternshipContract(internshipContract)
}

func (s *Service) DeleteInternshipContract(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteInternshipContract(id, deletedBy)
}

//internship Document Services

func (s *Service) CreateInternshipDocument(internshipDocument *InternshipDocument) error {
	if internshipDocument.DocumentFile == "" {
		return errors.New("insert document is required")
	}

	if internshipDocument.Description == "" {
		return errors.New("fill document description is required")
	}

	if !documentType[internshipDocument.Document] {
		return errors.New("invalid document type")
	}

	return s.Repo.CreateInternshipDocument(internshipDocument)
}

func (s *Service) GetInternshipDocument() ([]InternshipDocument, error) {
	return s.Repo.GetAllInternshipDocument()
}

func (s *Service) GetInternshipDocumentByID(id uint64) (*InternshipDocument, error) {
	return s.Repo.GetInternshipDocumentByID(id)
}

func (s *Service) UpdateInternshipDocument(internshipDocument *InternshipDocument) error {
	return s.Repo.UpdateInternshipDocument(internshipDocument)
}

func (s *Service) DeleteInternshipDocument(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteInternshipDocument(id, deletedBy)
}

//internship Evaluation Services

func (s *Service) CreateInternshipEvaluation(internshipEvaluation *InternshipEvaluation) error {
	if !evaluationRating[internshipEvaluation.Rating] {
		return errors.New("invalid Rating")
	}

	return s.Repo.CreateInternshipEvaluation(internshipEvaluation)
}

func (s *Service) GetInternshipEvaluation() ([]InternshipEvaluation, error) {
	return s.Repo.GetAllInternshipEvaluation()
}

func (s *Service) GetInternshipEvaluationByID(id uint64) (*InternshipEvaluation, error) {
	return s.Repo.GetInternshipEvaluationByID(id)
}

func (s *Service) UpdateInternshipEvaluation(internshipEvaluation *InternshipEvaluation) error {
	return s.Repo.UpdateInternshipEvaluation(internshipEvaluation)
}

func (s *Service) DeleteInternshipEvaluation(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteInternshipEvaluation(id, deletedBy)
}

//internship Progress Services

func (s *Service) CreateInternshipProgress(internshipProgress *InternshipProgress) error {
	if internshipProgress.Date == nil {
		return errors.New("Date is required")
	}

	if internshipProgress.DescriptionProgress == "" {
		return errors.New("Description is required")
	}

	if !progressStatus[internshipProgress.Status] {
		return errors.New("invalid progress status")
	}

	return s.Repo.CreateInternshipProgress(internshipProgress)
}

func (s *Service) GetInternshipProgress() ([]InternshipProgress, error) {
	return s.Repo.GetAllInternshipProgress()
}

func (s *Service) GetInternshipProgressByID(id uint64) (*InternshipProgress, error) {
	return s.Repo.GetInternshipProgressByID(id)
}

func (s *Service) UpdateInternshipProgress(internshipProgress *InternshipProgress) error {
	return s.Repo.UpdateInternshipProgress(internshipProgress)
}

func (s *Service) DeleteInternshipProgress(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteInternshipProgress(id, deletedBy)
}
