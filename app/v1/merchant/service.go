package merchant

type Service interface {
	CreateMerchant(inputData CreateMerchantInput) (Merchant, error)
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

	// create data nya
	newMerchant, err := s.repository.Save(merchant)
	if err != nil {
		return newMerchant, err
	}

	return newMerchant, nil
}
