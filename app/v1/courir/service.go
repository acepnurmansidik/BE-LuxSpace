package courir

type Service interface {
	GetCourirs() ([]Courir, error)
	GetDetailCourir(inputID CourirDetailInput) (Courir, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCourirs() ([]Courir, error) {
	courirs, err := s.repository.FindAll()

	if err != nil {
		return courirs, err
	}

	return courirs, nil
}

func (s *service) GetDetailCourir(inputID CourirDetailInput) (Courir, error) {
	courir, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return courir, err
	}

	return courir, nil
}
