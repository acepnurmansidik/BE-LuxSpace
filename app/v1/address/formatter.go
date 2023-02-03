package address

type FormatAddress struct {
	ID          int    `json:"id"`
	AddressName string `json:"address_name"`
	IsPrimary   int    `json:"is_primary"`
	OwnerName   string `json:"owner_name"`
}

func FormatterAddress(address Address) FormatAddress {
	formatter := FormatAddress{}
	formatter.ID = address.ID
	formatter.AddressName = address.AddressName
	formatter.IsPrimary = address.IsPrimary
	formatter.OwnerName = address.OwnerName

	return formatter
}

func FormatterAddressList(addressList []Address) []FormatAddress {
	var formatter []FormatAddress

	for _, address := range addressList {
		formatter = append(formatter, FormatterAddress(address))
	}

	return formatter
}
