package user

type User struct {
	ID          int
	Email       string
	Password    string
	Username    string
	Avatar      string
	PhoneNumber int
	Role        string
	CodeOtp     string
	IsActive    int
}
