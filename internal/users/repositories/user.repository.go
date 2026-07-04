package user_repositories

import (
	"github.com/go-minstack/go-minstack/repository"
	"gorm.io/gorm"
	user_entities "task-api/internal/users/entities"
)

type UserRepository struct {
	*repository.Repository[user_entities.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{repository.NewRepository[user_entities.User](db)}
}

func (r *UserRepository) FindByEmail(email string) (*user_entities.User, error) {
	return r.FindOne(repository.Where("email = ?", email))
}
