package merchant

type FormatMerchant struct {
	ID              int    `json:"id" binding:"required"`
	MerchantName    string `json:"merchant_name" binding:"required"`
	MerchantAddress string `json:"merchant_address" binding:"required"`
	Avatar          string `json:"image_url" binding:"required"`
}

func FormatterMerchant(merchant Merchant) FormatMerchant {
	formatter := FormatMerchant{}
	formatter.ID = merchant.ID
	formatter.MerchantName = merchant.MerchantName
	formatter.MerchantAddress = merchant.MerchantAddress
	formatter.Avatar = merchant.Avatar

	return formatter
}
