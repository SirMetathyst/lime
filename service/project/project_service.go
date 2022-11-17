package project

import (
	"context"
	"errors"
	"fmt"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project Project) (string, error)
	ReadProject(ctx context.Context, id string) (Project, error)
	UpdateProject(ctx context.Context, project Project) error
	DeleteProject(ctx context.Context, projectID string) error
	ListProjects(ctx context.Context) ([]Project, error)
}

type projectService struct {
	projectRepository ProjectRepository
	TaskService
}

func NewProjectService(taskService TaskService, projectRepository ProjectRepository) (ProjectService, error) {
	if taskService == nil {
		return nil, errors.New("project-service: task service is nil")
	}
	if projectRepository == nil {
		return nil, errors.New("project-service: project repository is nil")
	}
	return &projectService{TaskService: taskService, projectRepository: projectRepository}, nil
}

func (s *projectService) CreateProject(ctx context.Context, project Project) (string, error) {

	id, err := s.projectRepository.CreateProject(ctx, project)
	if err != nil {
		return "", fmt.Errorf("project-service[create-project]: %w", err)
	}

	return id, nil
}

func (s *projectService) ReadProject(ctx context.Context, id string) (Project, error) {

	readProject, err := s.projectRepository.ReadProject(ctx, id)
	if err != nil {
		return Project{}, fmt.Errorf("project-service[read-project]: %w", err)
	}

	return readProject, nil
}

func (s *projectService) UpdateProject(ctx context.Context, project Project) error {

	err := s.projectRepository.UpdateProject(ctx, project)
	if err != nil {
		return fmt.Errorf("project-service[update-project]: %w", err)
	}

	return nil
}

func (s *projectService) DeleteProject(ctx context.Context, projectID string) error {

	err := s.projectRepository.DeleteProject(ctx, projectID)
	if err != nil {
		return fmt.Errorf("project-service[delete-project]: %w", err)
	}

	err = s.TaskService.DeleteTasks(ctx, projectID)
	if err != nil {
		return fmt.Errorf("project-service[delete-project]: %w", err)
	}

	return nil
}

func (s *projectService) ListProjects(ctx context.Context) ([]Project, error) {

	projects, err := s.projectRepository.ListProjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("project-service[list-projects]: %w", err)
	}

	return projects, nil
}
