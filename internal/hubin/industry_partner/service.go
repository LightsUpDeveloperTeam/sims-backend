package industry_partner

import "errors"

var validCollaborationStatus = map[string]bool{
	"Active":   true,
	"Inactive": true,
}

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

//Collaboration History Services

func (s *Service) CreateCollaborationHistory(collaborationHistory *CollaborationHistory) error {
	if collaborationHistory.Description == "" {
		return errors.New("Fill the description!")
	}

	if !validCollaborationStatus[collaborationHistory.Status] {
		return errors.New("invalid collaboration status")
	}

	return s.Repo.CreateCollaborationHistory(collaborationHistory)
}

func (s *Service) GetCollaborationHistory() ([]CollaborationHistory, error) {
	return s.Repo.GetCollaborationHistory()
}

func (s *Service) GetCollaborationHistoryByID(id uint64) (*CollaborationHistory, error) {
	return s.Repo.GetCollaborationHistoryByID(id)
}

func (s *Service) UpdateCollaborationHistory(collaborationHistory *CollaborationHistory) error {
	return s.Repo.UpdateCollaborationHistory(collaborationHistory)
}

func (s *Service) DeleteCollaborationHistory(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteCollaborationHistory(id, deletedBy)
}

//Memorandum Of Understanding Services

func (s *Service) CreateMemorandumOfUnderstanding(memorandumOfUnderstanding *MemorandumOfUnderstanding) error {
	if memorandumOfUnderstanding.MOUNumber == "" {
		return errors.New("Insert the MoU number!")
	}

	if memorandumOfUnderstanding.Description == "" {
		return errors.New("Fill the description!")
	}

	return s.Repo.CreateMemorandumOfUnderstanding(memorandumOfUnderstanding)
}

func (s *Service) GetMemorandumOfUnderstanding() ([]MemorandumOfUnderstanding, error) {
	return s.Repo.GetAllMemorandumOfUnderstanding()
}

func (s *Service) GetMemorandumOfUnderstandingByID(id uint64) (*MemorandumOfUnderstanding, error) {
	return s.Repo.GetMemorandumOfUnderstandingByID(id)
}

func (s *Service) UpdateMemorandumOfUnderstanding(memorandumOfUnderstanding *MemorandumOfUnderstanding) error {
	return s.Repo.UpdateMemorandumOfUnderstanding(memorandumOfUnderstanding)
}

func (s *Service) DeleteMemorandumOfUnderstanding(id uint64, deletedBy uint64) error {
	return s.Repo.DeleteMemorandumOfUnderstanding(id, deletedBy)
}
