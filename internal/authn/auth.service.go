package authn

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-minstack/auth"
	"golang.org/x/crypto/bcrypt"
	"task-api/internal/users/dto"
	user_repos "task-api/internal/users/repositories"
)

type AuthService struct {
	users *user_repos.UserRepository
	jwt   *auth.JwtService
	log   *slog.Logger
}

func NewAuthService(users *user_repos.UserRepository, jwt *auth.JwtService, log *slog.Logger) *AuthService {
	return &AuthService{users: users, jwt: jwt, log: log}
}

func (s *AuthService) Login(input user_dto.LoginDto) (string, error) {
	user, err := s.users.FindByEmail(input.Email)
	if err != nil {
		s.log.Warn("login failed: user not found")
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		s.log.Warn("login failed: wrong password", "user_id", user.ID)
		return "", errors.New("invalid credentials")
	}

	token, err := s.jwt.Sign(auth.Claims{
		Subject: fmt.Sprintf("%d", user.ID),
		Name:    user.Name,
	}, 24*time.Hour)
	if err != nil {
		return "", err
	}

	s.log.Info("user authenticated", "user_id", user.ID)
	return token, nil
}
