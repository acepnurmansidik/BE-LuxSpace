package courir

type Service interface {
	GetCourirs() ([]Courir, error)
	GetDetailCourir(inputID CourirDetailInput) (Courir, error)
	CreateCourir(inputData CreateCourirInput) (Courir, error)
	UpdateCourir(inputID CourirDetailInput, inputData CreateCourirInput) (Courir, error)
	DeleteCourir(inputID CourirDetailInput) (Courir, error)
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

func (s *service) CreateCourir(inputData CreateCourirInput) (Courir, error) {
	// mapping data dari inputan ke dalam struct
	courir := Courir{}
	courir.Name = inputData.Name

	newCourir, err := s.repository.Save(courir)
	if err != nil {
		return newCourir, err
	}
	return newCourir, nil
}

func (s *service) UpdateCourir(inputID CourirDetailInput, inputData CreateCourirInput) (Courir, error) {
	// cari courir berdasarkan id nya
	courir, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return courir, err
	}
	// mapping data courir
	courir.Name = inputData.Name
	// update data
	newCourir, err := s.repository.Update(courir)
	if err != nil {
		return newCourir, err
	}

	return newCourir, nil
}

func (s *service) DeleteCourir(inputID CourirDetailInput) (Courir, error) {
	// cari id kurir
	courir, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return courir, err
	}

	// delete kurirnya
	newCourir, err := s.repository.Destroy(courir)
	if err != nil {
		return courir, err
	}

	return newCourir, nil
}
