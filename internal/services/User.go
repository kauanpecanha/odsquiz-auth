package services

import (
	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	return s.Repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.ReadUsers()
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.Repo.ReadUserByID(id)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	return s.Repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.Repo.DeleteUser(id)
}