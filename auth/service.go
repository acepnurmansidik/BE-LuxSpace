package auth

import (
	"LuxSpace/app/v1/user"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(inputData user.User) (string, error)
}

type jwtService struct{}

// agar semua kontrak bisa ke import
func NewService() *jwtService {
	return &jwtService{}
}

var SECRET_KEY = []byte("LUXSPACE_PROJECT")

func (s *jwtService) GenerateToken(inputData user.User) (string, error) {
	// buat payload untuk jwt nya
	payload := jwt.MapClaims{}
	// mapping data ke dalam payload
	payload["user_id"] = inputData.ID
	payload["email"] = inputData.Email
	payload["role"] = inputData.Role
	payload["username"] = inputData.Username

	// masukan payload beserta algoritma yang digunakan
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// buat verify signature/ secret key
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
