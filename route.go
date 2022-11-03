package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route(r *chi.Mux, service Service) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := RenderLandingPage(w); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// LIST PROJECTS PAGE
	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		projects, _ := service.ListProjects(r.Context())

		data := ListProjectPage{Projects: projects}
		if err := RenderListProjectPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// NEW PROJECT PAGE
	r.Get("/dashboard/new", func(w http.ResponseWriter, r *http.Request) {

		if err := RenderNewProjectPage(w, NewProjectPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// SUBMIT NEW PROJECT
	r.Post("/dashboard/new", func(w http.ResponseWriter, r *http.Request) {

		// TODO: Error checking
		r.ParseForm()
		formProjectTitle := r.Form.Get("project-title")

		// TODO: Sanitize user input
		newProject, _ := NewProject()
		newProject.Title = formProjectTitle

		// TODO: Return to same page with error message
		service.CreateProject(r.Context(), newProject)

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// LIST TASKS PAGE
	r.Get("/p/{ID}", func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "ID")

		tasks, _ := service.ListTasks(r.Context(), id)

		data := ListTaskPage{ID: id, Tasks: tasks}
		if err := RenderListTaskPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// NEW TASK PAGE
	r.Get("/p/{ID}/new", func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "ID")

		data := NewTaskPage{ID: id}

		if err := RenderNewTaskPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// SUBMIT NEW TASK
	r.Post("/p/{ID}/new", func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "ID")

		// TODO: Error checking
		r.ParseForm()
		formTaskTitle := r.Form.Get("task-title")

		// TODO: Sanitize user input
		newTask, _ := NewTask()
		newTask.Title = formTaskTitle

		// TODO: Return to same page with error message
		service.CreateTask(r.Context(), id, newTask)

		http.Redirect(w, r, fmt.Sprintf("/p/%s", id), http.StatusSeeOther)
	})

	//r.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(StaticFS))).ServeHTTP)
	r.HandleFunc("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("."))).ServeHTTP)
}
