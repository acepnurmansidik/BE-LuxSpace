package address

type AddressUserInput struct {
	ID int
}

type CreateAddressInput struct {
	AddressName string `json:"address_name" binding:"required"`
	IsPrimary   string `json:"is_primary"`
	OwnerName   string `json:"owner_name" binding:"required"`
	UserId      int    `json:"user_id"`
}
