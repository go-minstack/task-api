package authn

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-minstack/auth"
	"golang.org/x/crypto/bcrypt"
	"task-api/internal/users/dto"
	user_repos "task-api/internal/users/repositories"
)

type AuthService struct {
	users *user_repos.UserRepository
	jwt   *auth.JwtService
}

func NewAuthService(users *user_repos.UserRepository, jwt *auth.JwtService) *AuthService {
	return &AuthService{users: users, jwt: jwt}
}

func (s *AuthService) Login(input user_dto.LoginDto) (string, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return s.jwt.Sign(auth.Claims{
		Subject: fmt.Sprintf("%d", user.ID),
		Name:    user.Name,
	}, 24*time.Hour)
}
