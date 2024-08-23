package fundwave

type Service interface {
	Findfundawave(userID int) ([]Fundwave, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Findfundawave(userID int) ([]Fundwave, error) {
	if userID != 0 {
		fundwave, err := s.repository.FindByUserID(userID)
		if err != nil {
			return fundwave, err
		}
		return fundwave, nil
	}

	fundwave, err := s.repository.FindAll()
	if err != nil {
		return fundwave, err
	}
	return fundwave, nil
}
