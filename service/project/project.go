package project

import "fmt"

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
