package user

type FormatUserLogin struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FormatterUserLogin(user User) FormatUserLogin {
	// mapping data responnya
	formatter := FormatUserLogin{}
	formatter.Username = user.Username
	formatter.Email = user.Email

	return formatter
}
