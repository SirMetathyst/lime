package main

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

type Project struct {
	ID    string
	Title string

	Tasks []Task
}

func NewProject() (Project, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return Project{}, err
	}

	return Project{ID: id.String()}, nil
}

func findTaskWithID(tasks []Task, id string) (int, Task, error) {
	for index, task := range tasks {
		if task.ID == id {
			return index, task, nil
		}
	}
	return -1, Task{}, errors.New("service: task not found")
}

type Task struct {
	ID    string
	Title string
}

func NewTask() (Task, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return Task{}, err
	}

	return Task{ID: id.String()}, nil
}

type Service interface {
	CreateProject(ctx context.Context, project Project) error
	ListProjects(ctx context.Context) ([]Project, error)

	CreateTask(ctx context.Context, projectID string, task Task) error
	ListTasks(ctx context.Context, projectID string) ([]Task, error)
}

type service struct {
	projects map[string]*Project
}

func NewService() Service {
	return &service{projects: map[string]*Project{}}
}

func (s *service) CreateProject(ctx context.Context, project Project) error {

	if _, ok := s.projects[project.ID]; ok {
		return errors.New("service: project found")
	}

	s.projects[project.ID] = &project

	return nil
}

func (s *service) ListProjects(ctx context.Context) ([]Project, error) {

	var projects []Project

	for _, project := range s.projects {
		projects = append(projects, *project)
	}

	return projects, nil
}

func (s *service) CreateTask(ctx context.Context, projectID string, task Task) error {

	if _, ok := s.projects[projectID]; !ok {
		return errors.New("service: project not found")
	}

	_, _, err := findTaskWithID(s.projects[projectID].Tasks, task.ID)
	if err == nil {
		return errors.New("service: task found")
	}

	project := s.projects[projectID]
	project.Tasks = append(s.projects[projectID].Tasks, task)

	s.projects[projectID] = project

	return nil
}

func (s *service) ListTasks(ctx context.Context, projectID string) ([]Task, error) {

	if _, ok := s.projects[projectID]; !ok {
		return nil, errors.New("service: project not found")
	}

	var tasks []Task
	for _, task := range s.projects[projectID].Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
