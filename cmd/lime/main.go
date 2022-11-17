package main

// DUE DATE: 16th Nov, HALF POINT: 9th Nov

// Feature List
// - [ ] User AuthN
// - [ ] User AuthZ
// - [ ] User Login/Registration
// - [ ] Landing Page
// - [ ] ReadProject CRUD
// - [ ] ReadTask CRUD
// - [ ] CSRF
// - [ ] Session

import (
	"github.com/SirMetathyst/lime/service/project"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/log"
	"os"
)

func main() {

	// LOGGING
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// TASK REPOSITORY
	taskRepository := project.Must(project.NewTaskRepository())

	// PROJECT REPOSITORY
	projectRepository := project.Must(project.NewProjectRepository())

	// TASK SERVICE
	var taskService project.TaskService
	taskService = project.Must(project.NewTaskService(taskRepository))

	// PROJECT SERVICE
	var projectService project.ProjectService
	projectService = project.Must(project.NewProjectService(taskService, projectRepository))

	// SERVICE
	var service project.Service
	service = project.Must(project.NewService(projectService, taskService))
	service = project.Must(project.NewServiceLoggingMiddleware(log.With(logger, "component", "service"), service))

	// ROUTING
	r := chi.NewRouter()
	r.Use(project.NewHTTPLoggingMiddleware(logger))
	r.Use(middleware.Recoverer)

	Route(r, service)
	Serve(r)

}
