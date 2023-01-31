package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(inputData CreateUserInput) (User, error)
	IsEmailAvailable(inputData CheckEmailInput) (bool, error)
	Login(inputData LoginInput) (User, error)
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

func (s *service) Login(inputData LoginInput) (User, error) {
	user, err := s.repository.FindByEmail(inputData.Email)
	if err != nil {
		return user, err
	}

	// cek jika emailnya tidak ada
	if user.ID == 0 {
		return user, errors.New("Email not registered")
	}

	// cek passwordnya
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputData.Password))
	if result != nil {
		return user, errors.New("Password no match")
	}

	return user, nil
}
