package user

type FormatUserRegister struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func FormatterUserRegister(user User) FormatUserRegister {
	// mapping data responnya
	formatter := FormatUserRegister{}
	formatter.Username = user.Username
	formatter.Email = user.Email

	return formatter
}

type FormatUserLogin struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FormatterUserLogin(user User, token string) FormatUserLogin {
	// mapping data responnya
	formatter := FormatUserLogin{}
	formatter.ID = user.ID
	formatter.Email = user.Email
	formatter.Role = user.Role
	formatter.Username = user.Username
	formatter.Token = token

	return formatter
}
