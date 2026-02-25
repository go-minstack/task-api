package tasks

import (
	"errors"
	"strconv"

	"github.com/go-minstack/auth"
	"task-api/internal/tasks/dto"
	task_entities "task-api/internal/tasks/entities"
	task_repos "task-api/internal/tasks/repositories"
)

type TaskService struct {
	tasks *task_repos.TaskRepository
}

func NewTaskService(tasks *task_repos.TaskRepository) *TaskService {
	return &TaskService{tasks: tasks}
}

func (s *TaskService) List(claims *auth.Claims) ([]dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	tasks, err := s.tasks.FindByUserID(uint(userID))
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.TaskDto, len(tasks))
	for i, t := range tasks {
		dtos[i] = dto.NewTaskDto(&t)
	}
	return dtos, nil
}

func (s *TaskService) Create(claims *auth.Claims, input dto.CreateTaskDto) (*dto.TaskDto, error) {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	task := &task_entities.Task{
		Title:       input.Title,
		Description: input.Description,
		UserID:      uint(userID),
	}
	if err := s.tasks.Create(task); err != nil {
		return nil, err
	}
	result := dto.NewTaskDto(task)
	return &result, nil
}

func (s *TaskService) Get(claims *auth.Claims, id uint) (*dto.TaskDto, error) {
	task, err := s.tasks.FindByID(id)
	if err != nil {
		return nil, err
	}
	if err := s.assertOwner(claims, task); err != nil {
		return nil, err
	}
	result := dto.NewTaskDto(task)
	return &result, nil
}

func (s *TaskService) Update(claims *auth.Claims, id uint, input dto.UpdateTaskDto) (*dto.TaskDto, error) {
	task, err := s.tasks.FindByID(id)
	if err != nil {
		return nil, err
	}
	if err := s.assertOwner(claims, task); err != nil {
		return nil, err
	}

	columns := map[string]interface{}{}
	if input.Title != "" {
		columns["title"] = input.Title
	}
	if input.Description != "" {
		columns["description"] = input.Description
	}
	if input.Done != nil {
		columns["done"] = *input.Done
	}
	if err := s.tasks.UpdatesByID(id, columns); err != nil {
		return nil, err
	}

	return s.Get(claims, id)
}

func (s *TaskService) Delete(claims *auth.Claims, id uint) error {
	task, err := s.tasks.FindByID(id)
	if err != nil {
		return err
	}
	if err := s.assertOwner(claims, task); err != nil {
		return err
	}
	return s.tasks.DeleteByID(id)
}

func (s *TaskService) assertOwner(claims *auth.Claims, task *task_entities.Task) error {
	userID, _ := strconv.ParseUint(claims.Subject, 10, 64)
	if task.UserID != uint(userID) {
		return errors.New("forbidden")
	}
	return nil
}
