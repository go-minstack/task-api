package users

import (
	"golang.org/x/crypto/bcrypt"
	"task-api/internal/users/dto"
	user_entities "task-api/internal/users/entities"
	user_repos "task-api/internal/users/repositories"
)

type UserService struct {
	users *user_repos.UserRepository
}

func NewUserService(users *user_repos.UserRepository) *UserService {
	return &UserService{users: users}
}

func (s *UserService) Register(input dto.RegisterDto) (*dto.UserDto, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &user_entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
	}
	if err := s.users.Create(user); err != nil {
		return nil, err
	}

	result := dto.NewUserDto(user)
	return &result, nil
}

func (s *UserService) Me(id uint) (*dto.UserDto, error) {
	user, err := s.users.FindByID(id)
	if err != nil {
		return nil, err
	}
	result := dto.NewUserDto(user)
	return &result, nil
}
