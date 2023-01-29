package category

type Service interface {
	GetDetailCategory(inputID CategoryDetailInput) (Category, error)
	GetCategories(Name string) ([]Category, error)
	CreateCategory(inputData CreateCategoryInput) (Category, error)
	UpdateCategory(inputID CategoryDetailInput, inputData CreateCategoryInput) (Category, error)
	DeleteCategory(inputID CategoryDetailInput) (Category, error)
}

type service struct {
	repository Repository
}

// CreateCategory implements Service
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCategories(Name string) ([]Category, error) {
	if len(Name) > 0 {
		category, err := s.repository.FindByName(Name)
		if err != nil {
			return category, err
		}

		return category, nil
	}

	category, err := s.repository.FindAll()
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) GetDetailCategory(inputID CategoryDetailInput) (Category, error) {
	category, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *service) CreateCategory(inputData CreateCategoryInput) (Category, error) {
	// mapping data dari inputan ke dalam struct
	category := Category{}
	category.Name = inputData.Name

	// simpan datanya
	newCategory, err := s.repository.Save(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) UpdateCategory(inputID CategoryDetailInput, inputData CreateCategoryInput) (Category, error) {
	// cari id category
	category, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return category, err
	}

	// mapping data yang diambil berdasarkan id
	category.Name = inputData.Name

	// lalu update datanya
	updateCategory, err := s.repository.Update(category)
	if err != nil {
		return updateCategory, err
	}

	return updateCategory, nil
}

func (s *service) DeleteCategory(inputID CategoryDetailInput) (Category, error) {
	// cek apakah datanya ada
	category, err := s.repository.FindByID(inputID.ID)
	// cek apakah ada error dan datanya ada
	if err != nil || category.ID == 0 {
		return category, err
	}

	newCategory, err := s.repository.Destroy(category)
	if err != nil {
		return newCategory, err
	}

	return category, nil
}
