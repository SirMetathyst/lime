package main

import (
	"context"
	"errors"
	"fmt"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, projectID string, task Task) (string, error)
	ReadTask(ctx context.Context, projectID string, taskID string) (Task, error)
	UpdateTask(ctx context.Context, projectID string, task Task) error
	DeleteTask(ctx context.Context, projectID string, taskID string) error
	DeleteTasks(ctx context.Context, projectID string) error
	ListTasks(ctx context.Context, projectID string) ([]Task, error)
}

type InMemoryMapTaskRepository struct {
	Task map[string][]Task
}

func NewInMemoryMapTaskRepository() (TaskRepository, error) {
	return &InMemoryMapTaskRepository{Task: map[string][]Task{}}, nil
}

func (s *InMemoryMapTaskRepository) CreateTask(ctx context.Context, projectID string, task Task) (string, error) {

	id, err := generateIDWhenNotGiven(task.ID)
	if err != nil {
		return "", fmt.Errorf("repository[create-task]: %w", err)
	}

	task.ID = id
	_, err = s.ReadTask(ctx, projectID, task.ID)
	if err == nil {
		return "", errors.New("repository[create-task]: task found")
	}

	s.Task[projectID] = append(s.Task[projectID], task)

	return id, nil
}

func (s *InMemoryMapTaskRepository) ReadTask(ctx context.Context, projectID string, taskID string) (Task, error) {

	for _, currentTask := range s.Task[projectID] {
		if currentTask.ID == taskID {
			return currentTask, nil
		}
	}

	return Task{}, errors.New("repository[read-task]: task not found")
}

func (s *InMemoryMapTaskRepository) UpdateTask(ctx context.Context, projectID string, task Task) error {

	if _, err := s.ReadTask(ctx, projectID, task.ID); err != nil {
		return errors.New("repository[update-task]: task not found")
	}

	for index, currentTask := range s.Task[projectID] {
		if currentTask.ID == task.ID {
			s.Task[projectID][index] = task
		}
	}

	return nil
}

func (s *InMemoryMapTaskRepository) DeleteTask(ctx context.Context, projectID string, taskID string) error {

	for index, task := range s.Task[projectID] {
		if task.ID == taskID {
			s.Task[projectID] = append(s.Task[projectID][:index], s.Task[projectID][index+1:]...)
			return nil
		}
	}

	return errors.New("repository[delete-task]: task not found")
}

func (s *InMemoryMapTaskRepository) DeleteTasks(ctx context.Context, projectID string) error {

	delete(s.Task, projectID)

	return nil
}

func (s *InMemoryMapTaskRepository) ListTasks(ctx context.Context, projectID string) (tasks []Task, err error) {

	for _, currentTask := range s.Task[projectID] {
		tasks = append(tasks, currentTask)
	}

	return tasks, nil
}
