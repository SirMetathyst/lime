package project

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

type projectRepository struct {
	project map[string]Project
}

func NewProjectRepository() (ProjectRepository, error) {
	return &projectRepository{project: map[string]Project{}}, nil
}

func (s *projectRepository) CreateProject(ctx context.Context, project Project) (string, error) {

	id, err := generateIDWhenNotGiven(project.ID)
	if err != nil {
		return "", fmt.Errorf("repository[create-project]: %w", err)
	}

	project.ID = id
	if _, err := s.ReadProject(ctx, project.ID); err == nil {
		return id, errors.New("repository[create-project]: project found")
	}

	s.project[project.ID] = project

	return id, nil
}

func (s *projectRepository) ReadProject(ctx context.Context, projectID string) (Project, error) {

	if project, ok := s.project[projectID]; ok {
		return project, nil
	}

	return Project{}, errors.New("repository[project]: project not found")
}

func (s *projectRepository) UpdateProject(ctx context.Context, project Project) error {

	if _, ok := s.project[project.ID]; ok {
		s.project[project.ID] = project
		return nil
	}

	return errors.New("repository[update-project]: project not found")
}

func (s *projectRepository) DeleteProject(ctx context.Context, projectID string) error {

	delete(s.project, projectID)

	return nil
}

func (s *projectRepository) ListProjects(ctx context.Context) (projects []Project, err error) {

	for _, v := range s.project {
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
