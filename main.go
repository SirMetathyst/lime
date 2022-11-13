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
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/log"
	stdlog "log"
	"os"
	"time"
)

func main() {

	// LOGGING
	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	// TASK REPOSITORY
	taskRepository, err := NewInMemoryMapTaskRepository()
	if err != nil {
		stdlog.Fatalln(err)
	}

	// PROJECT REPOSITORY
	projectRepository, err := NewInMemoryMapProjectRepository()
	if err != nil {
		stdlog.Fatalln(err)
	}

	// TASK SERVICE
	var taskService TaskService
	taskService = &StdTaskService{taskRepository}

	// PROJECT SERVICE
	var projectService ProjectService
	projectService = &StdProjectService{ProjectRepository: projectRepository, TaskService: taskService}

	// STD SERVICE
	var service Service
	service = &StdService{ProjectService: projectService, TaskService: taskService}
	service = &StdServiceLoggingMiddleware{Logger: log.With(logger, "component", "service"), Next: service}

	// TEST DATA
	//CreateTestData(service)

	// ROUTING
	r := chi.NewRouter()
	r.Use(HTTPLoggingMiddleware(logger))
	r.Use(middleware.Recoverer)

	Route(r, service)
	Serve(r)

}

func CreateTestData(service Service) {

	project, _ := NewProject()
	project.Title = "Untitled Test ReadProject"

	projectID, err := service.CreateProject(context.Background(), project)
	if err != nil {
		stdlog.Println(err)
	}
	for i := 1; i < 200; i++ {

		task, _ := NewTaskWithID()
		task.Status = StatusDoing.String()
		task.Priority = DegreeHigh.String()
		task.Importance = DegreeHigh.String()
		task.DateCreated = time.Now()
		task.Title = fmt.Sprintf("Test ReadTask %d", i)

		if _, err := service.CreateTask(context.Background(), projectID, task); err != nil {
			stdlog.Println(err)
		}
	}
}
