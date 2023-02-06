package merchant

import (
	"errors"
)

type Service interface {
	CreateMerchant(inputData CreateMerchantInput) (Merchant, error)
	UpdateMerchant(inputData Merchant) (Merchant, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateMerchant(inputData CreateMerchantInput) (Merchant, error) {
	merchant := Merchant{}
	merchant.MerchantAddress = inputData.MerchantAddress
	merchant.MerchantName = inputData.MerchantName
	merchant.UserId = inputData.UserId

	// cari nama merchant nya(karena harus unik)
	nameIsExist, err := s.repository.FindByName(inputData.MerchantName)
	if err != nil {
		return nameIsExist, err
	}

	// validasi nama merchant jika sudah terdaftar
	if nameIsExist.MerchantName != "" {
		return nameIsExist, errors.New("name merchant has been register")
	}

	// create data nya
	newMerchant, err := s.repository.Save(merchant)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}

func (s *service) UpdateMerchant(inputData Merchant) (Merchant, error) {
	merchant, err := s.repository.Update(inputData)
	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
