package product

type Service interface {
	CreateProduct(inputDataProduct CreateProductInput) (Product, error)
	SaveUploadProductImages(inputImage CreateProductImagesInput) (bool, error)
	GetAllMerchantProduct(merchantID int) ([]Product, error)
	DeleteProduct(productID ProductDetailInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateProduct(inputDataProduct CreateProductInput) (Product, error) {
	// mapping datanya
	formatProduct := Product{}
	formatProduct.Title = inputDataProduct.Title
	formatProduct.Description = inputDataProduct.Description
	formatProduct.Price = inputDataProduct.Price
	formatProduct.Stock = inputDataProduct.Stock
	formatProduct.Weight = inputDataProduct.Weight
	formatProduct.CategoryId = inputDataProduct.CategoryID
	formatProduct.MerchantId = inputDataProduct.MerchantID

	newProduct, err := s.repository.SaveProduct(formatProduct)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) SaveUploadProductImages(inputImage CreateProductImagesInput) (bool, error) {
	// mapping datanta
	productImage := ProductImages{}
	productImage.IsDelete = inputImage.IsDelete
	productImage.Name = inputImage.Name
	productImage.ProductId = inputImage.ProductId

	// simpan ke databse
	_, err := s.repository.SaveProductImage(productImage)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GetAllMerchantProduct(merchantID int) ([]Product, error) {
	// cari merchant berdasarkan user yang sudah mendafatar
	products, err := s.repository.FindAllByMerchant(merchantID)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) DeleteProduct(productID ProductDetailInput) (Product, error) {
	// cari product berdasarkan id
	getProduct, err := s.repository.FindByID(productID.ID)
	if err != nil {
		return getProduct, err
	}
	// hapus productnya
	newProduct, err := s.repository.Destroy(getProduct)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}
