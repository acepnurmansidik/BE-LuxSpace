package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(inputData CreateUserInput) (User, error)
	IsEmailAvailable(inputData CheckEmailInput) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(inputData CreateUserInput) (User, error) {
	user := User{}
	// hashing password sebelum di simpan ke database
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(inputData.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	// mapping datanya
	user.Username = inputData.Username
	user.Email = inputData.Email
	user.Password = string(passwordHash)
	user.Role = "customer"

	// lalu simpan ke dalam database
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) IsEmailAvailable(inputData CheckEmailInput) (bool, error) {
	// cek jika email sebelumnya sudah terdaftar
	user, err := s.repository.FindByEmail(inputData.Email)
	if err != nil {
		return false, err
	}

	// cek jika emailnya terdaftar
	if user.ID != 0 {
		return true, nil
	}

	return false, nil
}
