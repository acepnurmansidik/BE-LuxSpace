package merchant

type MerchantDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateMerchantInput struct {
	MerchantName    string `form:"merchant_name" binding:"required"`
	MerchantAddress string `form:"merchant_address" binding:"required"`
	UserId          int    `form:"user_id"`
}

type CreateImageMerchant struct {
	Avatar string `form:"avatar"`
}
