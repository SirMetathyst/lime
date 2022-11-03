package main

// DUE DATE: 16th Nov, HALF POINT: 9th Nov

// Feature List
// - User AuthN
// - User AuthZ
// - User Login/Registration
// - Landing Page
// - Project CRUD
// - Task CRUD

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
)

func main() {

	// SERVICES
	var service Service
	service = NewService()

	// TEST DATA
	CreateTestData(service)

	// ROUTING
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	Route(r, service)
	Serve(r)

}

func CreateTestData(service Service) {

	project, _ := NewProject()
	project.Title = "Untitled Test Project"

	if err := service.CreateProject(context.Background(), project); err != nil {
		log.Println(err)
	}

	for i := 1; i < 5; i++ {
		task, _ := NewTask()
		task.Title = fmt.Sprintf("Test Task %d", i)

		if err := service.CreateTask(context.Background(), project.ID, task); err != nil {
			log.Println(err)
		}
	}
}
