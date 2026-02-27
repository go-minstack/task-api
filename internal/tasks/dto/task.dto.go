package task_dto

import task_entities "task-api/internal/tasks/entities"

type TaskDto struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID      uint   `json:"user_id"`
}

func NewTaskDto(t *task_entities.Task) TaskDto {
	return TaskDto{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Done:        t.Done,
		UserID:      t.UserID,
	}
}
