package main

import (
	"context"
	"fmt"
)

// todo: make task service part of project service?
type TaskService interface {
	CreateTask(ctx context.Context, projectID string, task Task) (string, error)
	ReadTask(ctx context.Context, projectID string, taskID string) (Task, error)
	UpdateTask(ctx context.Context, projectID string, task Task) error
	DeleteTask(ctx context.Context, projectID string, taskID string) error
	DeleteTasks(ctx context.Context, projectID string) error
	ListTasks(ctx context.Context, projectID string) ([]Task, error)
}

type StdTaskService struct {
	TaskRepository TaskRepository
}

func (s *StdTaskService) CreateTask(ctx context.Context, projectID string, task Task) (string, error) {

	if err := TaskValid(task); err != nil {
		return "", fmt.Errorf("service[create-task]: %w", err)
	}

	id, err := s.TaskRepository.CreateTask(ctx, projectID, task)
	if err != nil {
		return "", fmt.Errorf("task-service[create-task]: %w", err)
	}

	return id, nil
}

func (s *StdTaskService) ReadTask(ctx context.Context, projectID string, taskID string) (Task, error) {

	task, err := s.TaskRepository.ReadTask(ctx, projectID, taskID)
	if err != nil {
		return Task{}, fmt.Errorf("task-service[read-task]: %w", err)
	}

	return task, nil
}

func (s *StdTaskService) UpdateTask(ctx context.Context, projectID string, task Task) error {

	if err := TaskValid(task); err != nil {
		return fmt.Errorf("service[update-task]: %w", err)
	}

	err := s.TaskRepository.UpdateTask(ctx, projectID, task)
	if err != nil {
		return fmt.Errorf("task-service[update-task]: %w", err)
	}

	return nil
}

func (s *StdTaskService) DeleteTask(ctx context.Context, projectID string, taskID string) error {

	err := s.TaskRepository.DeleteTask(ctx, projectID, taskID)
	if err != nil {
		return fmt.Errorf("task-service[delete-task]: %w", err)
	}

	return nil
}

func (s *StdTaskService) DeleteTasks(ctx context.Context, projectID string) error {

	err := s.TaskRepository.DeleteTasks(ctx, projectID)
	if err != nil {
		return fmt.Errorf("task-service[delete-tasks]: %w", err)
	}

	return nil
}

func (s *StdTaskService) ListTasks(ctx context.Context, projectID string) ([]Task, error) {

	tasks, err := s.TaskRepository.ListTasks(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("task-service[list-tasks]: %w", err)
	}

	return tasks, nil
}
