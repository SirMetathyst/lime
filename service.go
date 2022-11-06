package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

type Project struct {
	ID    string
	Title string

	Tasks []Task
}

func NewProjectWithID() (Project, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return Project{}, fmt.Errorf("new-project: %w", err)
	}

	return Project{ID: id.String()}, nil
}

func findTaskWithID(tasks []Task, id string) (int, Task, error) {

	for index, task := range tasks {
		if task.ID == id {
			return index, task, nil
		}
	}

	return -1, Task{}, errors.New("find-task-with-id: task not found")
}

type Status string

var (
	StatusDontDo Status = "dont-do"
	StatusTodo   Status = "todo"
	StatusDoing  Status = "doing"
	StatusDone   Status = "done"
)

func (s Status) String() string {
	return string(s)
}

func StatusValid(v string) error {

	vs := Status(v)
	if vs == StatusDontDo || vs == StatusTodo || vs == StatusDoing || vs == StatusDone {
		return nil
	}

	return errors.New("status-valid: invalid")
}

type Degree string

var (
	DegreeLow    Degree = "low"
	DegreeMedium Degree = "medium"
	DegreeHigh   Degree = "high"
)

func (s Degree) String() string {
	return string(s)
}

func DegreeValid(v string) error {

	vs := Degree(v)
	if vs == DegreeLow || vs == DegreeMedium || vs == DegreeHigh {
		return nil
	}

	return errors.New("degree-valid: invalid")
}

func TitleValid(v string) error {

	if len(strings.TrimSpace(v)) == 0 {
		return errors.New("title-valid: is empty")
	}

	if len(v) > 200 {
		return errors.New("title-valid: length over 200")
	}

	return nil
}

type Task struct {
	ID          string
	Title       string
	Status      string
	Priority    string
	Importance  string
	DateCreated time.Time
}

func TaskValid(v Task) error {

	if v.ID == "" || strings.ContainsAny(v.ID, " \t\n\r") {
		return errors.New("task-valid: id invalid")
	}

	if err := TitleValid(v.Title); err != nil {
		return fmt.Errorf("task-valid: %w", err)
	}

	if err := StatusValid(v.Status); err != nil {
		return fmt.Errorf("task-valid: %w", err)
	}

	if err := DegreeValid(v.Priority); err != nil {
		return fmt.Errorf("task-valid: priority: %w", err)
	}

	if err := DegreeValid(v.Importance); err != nil {
		return fmt.Errorf("task-valid: importance: %w", err)
	}

	if v.DateCreated.IsZero() {
		return errors.New("task-valid: date created is the zero time instant")
	}

	return nil
}

func NewTaskWithID() (Task, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return Task{}, fmt.Errorf("new-task: %w", err)
	}

	return Task{ID: id.String()}, nil
}

// TODO: move to a displayTask struct with embedded task for the purpose of template displaying
func (s Task) DisplayDateCreated() string {
	return s.DateCreated.Format(time.RFC822)
}

type Service interface {
	CreateProject(ctx context.Context, project Project) error
	ListProjects(ctx context.Context) ([]Project, error)
	Project(ctx context.Context, id string) (Project, error)

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
		return errors.New("service[create-project]: project found")
	}

	s.projects[project.ID] = &project

	return nil
}

func (s *service) Project(ctx context.Context, id string) (Project, error) {

	for _, project := range s.projects {
		if project.ID == id {
			return *project, nil
		}
	}

	return Project{}, errors.New("service[project]: project not found")
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
		return errors.New("service[create-task]: project not found")
	}

	if err := TaskValid(task); err != nil {
		return fmt.Errorf("service[create-task]: %w", err)
	}

	_, _, err := findTaskWithID(s.projects[projectID].Tasks, task.ID)
	if err == nil {
		return errors.New("service[create-task]: task found")
	}

	project := s.projects[projectID]
	project.Tasks = append(s.projects[projectID].Tasks, task)

	s.projects[projectID] = project

	return nil
}

func (s *service) ListTasks(ctx context.Context, projectID string) ([]Task, error) {

	if _, ok := s.projects[projectID]; !ok {
		return nil, errors.New("service[list-tasks]: project not found")
	}

	var tasks []Task
	for _, task := range s.projects[projectID].Tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
