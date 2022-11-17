package project

import (
	"context"
	"net/http"
)

type TemplateService interface {

	// TASK

	NewTaskPage(ctx context.Context, w http.ResponseWriter, data NewTaskPage) error
	ListTask(ctx context.Context, w http.ResponseWriter, data ListTaskPage) error
	EditTask(ctx context.Context, w http.ResponseWriter, data EditTaskPage) error
	DeleteTask(ctx context.Context, w http.ResponseWriter, data DeleteTaskPage) error

	// PROJECT

	NewProject(ctx context.Context, w http.ResponseWriter, data NewProjectPage) error
	EditProject(ctx context.Context, w http.ResponseWriter, data EditProjectPage) error
}
