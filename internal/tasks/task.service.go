package tasks

import (
	"log/slog"
	"strconv"

	"github.com/go-minstack/auth"
	"task-api/internal/tasks/dto"
	task_entities "task-api/internal/tasks/entities"
	task_repos "task-api/internal/tasks/repositories"
)

type TaskService struct {
	tasks *task_repos.TaskRepository
	log   *slog.Logger
}

func NewTaskService(tasks *task_repos.TaskRepository, log *slog.Logger) *TaskService {
	return &TaskService{tasks: tasks, log: log}
}

func (s *TaskService) List(claims *auth.Claims) ([]task_dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	tasks, err := s.tasks.FindByUserID(uint(userID))
	if err != nil {
		s.log.Error("failed to list tasks", "user_id", userID, "error", err)
		return nil, err
	}
	s.log.Info("listed tasks", "user_id", userID, "count", len(tasks))
	dtos := make([]task_dto.TaskDto, len(tasks))
	for i, t := range tasks {
		dtos[i] = task_dto.NewTaskDto(&t)
	}
	return dtos, nil
}

func (s *TaskService) Create(claims *auth.Claims, input task_dto.CreateTaskDto) (*task_dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	task := &task_entities.Task{
		Title:       input.Title,
		Description: input.Description,
		UserID:      uint(userID),
	}
	if err := s.tasks.Create(task); err != nil {
		s.log.Error("failed to create task", "user_id", userID, "error", err)
		return nil, err
	}
	s.log.Info("task created", "task_id", task.ID, "user_id", userID)
	result := task_dto.NewTaskDto(task)
	return &result, nil
}

func (s *TaskService) Get(claims *auth.Claims, id uint) (*task_dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	task, err := s.tasks.FindByIDAndUserID(id, uint(userID))
	if err != nil {
		s.log.Error("task not found", "task_id", id, "user_id", userID)
		return nil, err
	}
	result := task_dto.NewTaskDto(task)
	return &result, nil
}

func (s *TaskService) Update(claims *auth.Claims, id uint, input task_dto.UpdateTaskDto) (*task_dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	task, err := s.tasks.FindByIDAndUserID(id, uint(userID))
	if err != nil {
		return nil, err
	}
	columns := map[string]interface{}{}
	if input.Title != "" {
		columns["title"] = input.Title
		task.Title = input.Title
	}
	if input.Description != "" {
		columns["description"] = input.Description
		task.Description = input.Description
	}
	if input.Done != nil {
		columns["done"] = *input.Done
		task.Done = *input.Done
	}
	if err := s.tasks.UpdatesByID(id, columns); err != nil {
		s.log.Error("failed to update task", "task_id", id, "error", err)
		return nil, err
	}
	s.log.Info("task updated", "task_id", id, "user_id", userID)
	result := task_dto.NewTaskDto(task)
	return &result, nil
}

func (s *TaskService) Delete(claims *auth.Claims, id uint) error {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	if _, err := s.tasks.FindByIDAndUserID(id, uint(userID)); err != nil {
		s.log.Error("task not found for deletion", "task_id", id, "user_id", userID)
		return err
	}
	if err := s.tasks.DeleteByID(id); err != nil {
		s.log.Error("failed to delete task", "task_id", id, "error", err)
		return err
	}
	s.log.Info("task deleted", "task_id", id, "user_id", userID)
	return nil
}
