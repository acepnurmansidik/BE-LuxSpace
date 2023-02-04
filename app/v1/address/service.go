package address

type Service interface {
	GetAllAddress(inputID AddressUserInput) ([]Address, error)
	GetDetailAddress(inputID AddressDetailInput) (Address, error)
	CreateAddress(inputData CreateAddressInput) (Address, error)
	UpdateAddress(inputID AddressDetailInput, inputData CreateAddressInput) (Address, error)
	DeleteAddress(inputID AddressDetailInput) (Address, error)
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

func (s *service) UpdateAddress(inputID AddressDetailInput, inputData CreateAddressInput) (Address, error) {
	// cari address user berdasarkan id nya
	address, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return address, err
	}
	// cek jika alamatnya di jadikan utama
	if inputData.IsPrimary == "true" {
		// update semuat alamatnya ke false kecuali yang ini
		addressUpdate, err := s.repository.FindAll(inputData.UserId)
		if err != nil {
			return address, err
		}

		for _, everyAddress := range addressUpdate {
			everyAddress.IsPrimary = "false"
			hashUpdate, err := s.repository.Update(everyAddress)
			if err != nil {
				return hashUpdate, err
			}
		}
	}

	// mapping data yang akan di updatente nya
	address.AddressName = inputData.AddressName
	address.OwnerName = inputData.OwnerName
	address.IsPrimary = inputData.IsPrimary

	// lakukan update
	addressUpdated, err := s.repository.Update(address)
	if err != nil {
		return addressUpdated, err
	}

	return addressUpdated, err
}

func (s *service) DeleteAddress(inputID AddressDetailInput) (Address, error) {
	// cari id yang akan di update
	address, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return address, err
	}

	// delete datanya
	_, err = s.repository.Destroy(address)
	if err != nil {
		return address, err
	}

	return address, nil
}

func (s *service) GetDetailAddress(inputID AddressDetailInput) (Address, error) {
	address, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return address, err
	}

	return address, nil
}
