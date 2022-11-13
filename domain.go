package main

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
	"time"
)

var (
	ErrTitleEmpty         = errors.New("title error: is empty")
	ErrTitleOverMaxLength = errors.New("title error: length over 200")
)

func NewID() (string, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("new-id: %w", err)
	}

	return id.String(), nil
}

type Project struct {
	ID    string
	Title string
}

func NewProjectWithID() (Project, error) {

	id, err := NewID()
	if err != nil {
		return Project{}, fmt.Errorf("new-project: %w", err)
	}

	return Project{ID: id}, nil
}

func NewProject() (Project, error) {
	return Project{}, nil
}

var (
	ErrNonExistentStatus = errors.New("status error: non existent status")
)

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

	return fmt.Errorf("status-valid: %w", ErrNonExistentStatus)
}

var (
	ErrNonExistentDegree = errors.New("degree error: non existent degree")
)

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

	return fmt.Errorf("degree-valid: %w", ErrNonExistentDegree)
}

func TitleValid(v string) error {

	if len(strings.TrimSpace(v)) == 0 {
		return ErrTitleEmpty
	}

	if len(v) > 200 {
		return ErrTitleOverMaxLength
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

var (
	ErrNonExistentPriority          = errors.New("priority error: non existent priority")
	ErrNonExistentImportance        = errors.New("importance error: non existent importance")
	ErrDateCreatedIsZeroTimeInstant = errors.New("date created error: date created is the zero time instant")
)

func TaskValid(v Task) error {

	//if v.ID == "" || strings.ContainsAny(v.ID, " \t\n\r") {
	//	return errors.New("task-valid: id invalid")
	//}

	if err := TitleValid(v.Title); err != nil {
		return fmt.Errorf("task-valid: %w", err)
	}

	if err := StatusValid(v.Status); err != nil {
		return fmt.Errorf("task-valid: %w", err)
	}

	if err := DegreeValid(v.Priority); err != nil {
		return fmt.Errorf("task-valid: %w", ErrNonExistentPriority)
	}

	if err := DegreeValid(v.Importance); err != nil {
		return fmt.Errorf("task-valid: %w", ErrNonExistentImportance)
	}

	if v.DateCreated.IsZero() {
		return fmt.Errorf("task-valid: %w", ErrDateCreatedIsZeroTimeInstant)
	}

	return nil
}

func NewTaskWithID() (Task, error) {

	id, err := NewID()
	if err != nil {
		return Task{}, fmt.Errorf("new-task: %w", err)
	}

	return Task{ID: id}, nil
}

func NewTask() (Task, error) {
	return Task{}, nil
}

// TODO: move to a displayTask struct with embedded task for the purpose of template displaying
func (s Task) DisplayDateCreated() string {
	return s.DateCreated.Format(time.RFC822)
}
