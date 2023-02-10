package merchant

import (
	"errors"
	"strconv"
)

type Service interface {
	CreateMerchant(inputData CreateMerchantInput) (Merchant, error)
	UpdateMerchant(fileLocation string, userID int) (Merchant, error)
	GetMerchantByUserID(userID int) (Merchant, error)
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

	// cek apakah user tersebut sadah punya merchant sebelumnya
	cekMerchant, err := s.repository.FindByUserID(inputData.UserId)
	if err != nil {
		return cekMerchant, err
	}

	// conversi string ke boolean
	isExits, _ := strconv.ParseBool(cekMerchant.MerchantName)
	// cek jika user sebelumnya sudah punya merchant
	if isExits {
		return cekMerchant, errors.New("you has merchant")
	}

	// create data nya
	newMerchant, err := s.repository.Save(merchant)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}

func (s *service) UpdateMerchant(fileLocation string, userID int) (Merchant, error) {
	merchant, err := s.repository.FindByUserID(userID)
	if err != nil {
		return merchant, err
	}

	// mapping data file location
	merchant.Avatar = fileLocation

	// update file location ke database
	newMerchant, err := s.repository.Update(merchant)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}

func (s *service) GetMerchantByUserID(userID int) (Merchant, error) {
	newMerchant, err := s.repository.FindByUserID(userID)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}
