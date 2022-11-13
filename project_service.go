package main

import (
	"context"
	"fmt"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project Project) (string, error)
	ReadProject(ctx context.Context, id string) (Project, error)
	UpdateProject(ctx context.Context, project Project) error
	DeleteProject(ctx context.Context, projectID string) error
	ListProjects(ctx context.Context) ([]Project, error)
}

type StdProjectService struct {
	ProjectRepository ProjectRepository
	TaskService
}

func (s *StdProjectService) CreateProject(ctx context.Context, project Project) (string, error) {

	id, err := s.ProjectRepository.CreateProject(ctx, project)
	if err != nil {
		return "", fmt.Errorf("project-service[create-project]: %w", err)
	}

	return id, nil
}

func (s *StdProjectService) ReadProject(ctx context.Context, id string) (Project, error) {

	project, err := s.ProjectRepository.ReadProject(ctx, id)
	if err != nil {
		return Project{}, fmt.Errorf("project-service[read-project]: %w", err)
	}

	return project, nil
}

func (s *StdProjectService) UpdateProject(ctx context.Context, project Project) error {

	err := s.ProjectRepository.UpdateProject(ctx, project)
	if err != nil {
		return fmt.Errorf("project-service[update-project]: %w", err)
	}

	return nil
}

func (s *StdProjectService) DeleteProject(ctx context.Context, projectID string) error {

	err := s.ProjectRepository.DeleteProject(ctx, projectID)
	if err != nil {
		return fmt.Errorf("project-service[delete-project]: %w", err)
	}

	err = s.TaskService.DeleteTasks(ctx, projectID)
	if err != nil {
		return fmt.Errorf("project-service[delete-project]: %w", err)
	}

	return nil
}

func (s *StdProjectService) ListProjects(ctx context.Context) ([]Project, error) {

	projects, err := s.ProjectRepository.ListProjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("project-service[list-projects]: %w", err)
	}

	return projects, nil
}
