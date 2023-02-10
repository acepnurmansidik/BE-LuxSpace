package product

import "LuxSpace/app/v1/merchant"

type Service interface {
	CreateProduct(inputDataProduct CreateProductInput) (Product, error)
	SaveUploadProductImages(inputImage CreateProductImagesInput) (bool, error)
	GetAllMerchantProduct(userID int) ([]Product, error)
}

type service struct {
	repository Repository
	merchant   merchant.Repository
}

func NewService(repository Repository, merchant merchant.Repository) *service {
	return &service{repository, merchant}
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
	productImage.IsPrimary = inputImage.IsPrimary
	productImage.Name = inputImage.Name
	productImage.ProductId = inputImage.ProductId

	// simpan ke databse
	_, err := s.repository.SaveProductImage(productImage)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GetAllMerchantProduct(userID int) ([]Product, error) {
	// cari merchant berdasarkan user yang sudah mendafatar
	merchantID, err := s.merchant.FindByUserID(userID)
	if err != nil {
		return []Product{}, err
	}

	// jika ada cari product nya
	products, err := s.repository.FindAllByMerchant(merchantID.ID)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}
