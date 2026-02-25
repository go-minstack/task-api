package dto

type UpdateTaskDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}
