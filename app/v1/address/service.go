package address

type Service interface {
	GetAllAddress(inputID AddressUserInput) ([]Address, error)
	CreateAddress(inputData CreateAddressInput) (Address, error)
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

func (s *service) CreateAddress(inputData CreateAddressInput) (Address, error) {
	address := Address{}
	// update semua is primary nya ke false
	addressWillUpdate, err := s.repository.FindAll(inputData.UserId)
	if err != nil {
		return address, err
	}

	// update alamat user ke false
	for _, everyAddress := range addressWillUpdate {
		everyAddress.IsPrimary = "false"
		// update alamat utamnya
		addressUpdated, err := s.repository.Update(everyAddress)
		if err != nil {
			return addressUpdated, err
		}
	}

	address.AddressName = inputData.AddressName
	address.OwnerName = inputData.OwnerName
	address.UserId = inputData.UserId
	address.IsPrimary = "true"

	newAddress, err := s.repository.Save(address)
	if err != nil {
		return newAddress, err
	}

	return newAddress, nil
}
