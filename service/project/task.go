package project

import (
	"fmt"
	"time"
)

type Task struct {
	ID          string
	Title       string
	Status      string
	Priority    string
	Importance  string
	DateCreated time.Time
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
