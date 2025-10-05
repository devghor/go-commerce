package service

import "go-commerce/internal/domain"

type UserService struct {
}

func (s *UserService) FindUserByEmail(email string) (*domain.User, error) {
	// Implement user registration logic here
	return nil, nil
}

func (s *UserService) SignUp(input any) error {
	// Implement user creation logic here
	return nil
}
