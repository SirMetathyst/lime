package project

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

var (
	ErrNonExistentPriority          = errors.New("priority error: non existent priority")
	ErrNonExistentImportance        = errors.New("importance error: non existent importance")
	ErrDateCreatedIsZeroTimeInstant = errors.New("date created error: date created is the zero time instant")
	ErrTitleEmpty                   = errors.New("title error: is empty")
	ErrTitleOverMaxLength           = errors.New("title error: length over 200")
	ErrNonExistentStatus            = errors.New("status error: non existent status")
	ErrNonExistentDegree            = errors.New("degree error: non existent degree")
)

func NewID() (string, error) {

	id, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("new-id: %w", err)
	}

	return id.String(), nil
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

	return fmt.Errorf("status-valid: %w", ErrNonExistentStatus)
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
