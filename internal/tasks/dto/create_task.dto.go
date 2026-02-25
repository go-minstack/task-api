package dto

type CreateTaskDto struct {
	Title       string `json:"title"       binding:"required"`
	Description string `json:"description"`
}
