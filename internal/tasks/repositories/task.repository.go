package task_repositories

import (
	"github.com/go-minstack/repository"
	task_entities "task-api/internal/tasks/entities"
	"gorm.io/gorm"
)

type TaskRepository struct {
	*repository.Repository[task_entities.Task]
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{repository.NewRepository[task_entities.Task](db)}
}

func (r *TaskRepository) FindByUserID(userID uint) ([]task_entities.Task, error) {
	return r.FindAll(repository.Where("user_id = ?", userID))
}
