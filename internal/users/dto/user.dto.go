package dto

import user_entities "task-api/internal/users/entities"

type UserDto struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDto(u *user_entities.User) UserDto {
	return UserDto{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
