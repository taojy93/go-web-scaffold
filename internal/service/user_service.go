package service

import (
	"go-web-scaffold/internal/models"
	"go-web-scaffold/internal/repository"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepository.DeleteUser(id)
}
