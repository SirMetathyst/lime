package main

import (
	"context"
	"errors"
	"fmt"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project Project) (string, error)
	ReadProject(ctx context.Context, projectID string) (Project, error)
	UpdateProject(ctx context.Context, project Project) error
	DeleteProject(ctx context.Context, projectID string) error
	ListProjects(ctx context.Context) ([]Project, error)
}

type StdInMemoryMapProjectRepository struct {
	Project map[string]Project
}

func NewInMemoryMapProjectRepository() (ProjectRepository, error) {
	return &StdInMemoryMapProjectRepository{Project: map[string]Project{}}, nil
}

func (s *StdInMemoryMapProjectRepository) CreateProject(ctx context.Context, project Project) (string, error) {

	id, err := generateIDWhenNotGiven(project.ID)
	if err != nil {
		return "", fmt.Errorf("repository[create-project]: %w", err)
	}

	project.ID = id
	if _, err := s.ReadProject(ctx, project.ID); err == nil {
		return id, errors.New("repository[create-project]: project found")
	}

	s.Project[project.ID] = project

	return id, nil
}

func (s *StdInMemoryMapProjectRepository) ReadProject(ctx context.Context, projectID string) (Project, error) {

	if project, ok := s.Project[projectID]; ok {
		return project, nil
	}

	return Project{}, errors.New("repository[project]: project not found")
}

func (s *StdInMemoryMapProjectRepository) UpdateProject(ctx context.Context, project Project) error {

	if _, ok := s.Project[project.ID]; ok {
		s.Project[project.ID] = project
		return nil
	}

	return errors.New("repository[update-project]: project not found")
}

func (s *StdInMemoryMapProjectRepository) DeleteProject(ctx context.Context, projectID string) error {

	delete(s.Project, projectID)

	return nil
}

func (s *StdInMemoryMapProjectRepository) ListProjects(ctx context.Context) (projects []Project, err error) {

	for _, v := range s.Project {
		projects = append(projects, v)
	}

	return projects, nil
}

func generateIDWhenNotGiven(id string) (string, error) {

	if len(id) == 0 {

		newID, err := NewID()
		if err != nil {
			return "", fmt.Errorf("generate-id-when-not-given: %w", err)
		}

		return newID, nil
	}

	return id, nil
}
