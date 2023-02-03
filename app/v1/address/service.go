package address

type Service interface {
	GetAllAddress(inputID AddressUserInput) ([]Address, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllAddress(inputID AddressUserInput) ([]Address, error) {
	address, err := s.repository.FindAll(inputID.ID)
	if err != nil {
		return address, err
	}

	return address, nil
}
