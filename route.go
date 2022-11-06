package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func Route(r *chi.Mux, service Service) {

	// APP...

	// LANDING PAGE

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := RenderLandingDesignPage(w); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// LOGIN

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := RenderLoginDesignPage(w); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// LIST PROJECTS PAGE
	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		data := DashboardPage{Projects: projects}
		if err := RenderDashboardPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// DASHBOARD POST FOR LOGIN TESTING
	r.Post("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// NEW PROJECT PAGE
	r.Get("/dashboard/new", func(w http.ResponseWriter, r *http.Request) {

		redirect := r.URL.Query().Get("redirect")

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		data := NewProjectPage{
			DashboardPage{
				Projects: projects,
				Page: Page{
					QueryParams: BuildQueryParams(redirect),
					Redirect:    redirect,
				},
			},
		}

		if err := RenderNewProjectPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// SUBMIT NEW PROJECT
	r.Post("/dashboard/new", func(w http.ResponseWriter, r *http.Request) {

		// todo: proper error handling
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		// todo: sanitize user input
		formProjectTitle := r.Form.Get("project-title")

		newProject, _ := NewProjectWithID()
		newProject.Title = formProjectTitle

		// todo: proper error handling
		if err := service.CreateProject(r.Context(), newProject); err != nil {
			log.Println(err)
		}

		// failure path:
		// ...
		// todo: return to same page with error message if an error occurred
		// todo: replace with new project error template or add error detail struct to template and conditionally check

		// success path:
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	// LIST TASKS PAGE
	r.Get("/dashboard/p/{ID}", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		selectedProject, err := service.Project(r.Context(), projectID)
		if err != nil {
			log.Println(err)
		}

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		tasks, err := service.ListTasks(r.Context(), projectID)
		if err != nil {
			log.Println(err)
		}

		data := ListTaskPage{
			ID:    projectID,
			Title: selectedProject.Title,
			DashboardPage: DashboardPage{
				Projects:        projects,
				SelectedProject: selectedProject,
				Page: Page{
					QueryParams: BuildQueryParams(fmt.Sprintf("/dashboard/p/%s", projectID)),
					Redirect:    fmt.Sprintf("/dashboard/p/%s", projectID),
				},
			},
			Tasks: tasks,
		}
		if err := RenderListTaskPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// NEW TASK PAGE
	r.Get("/dashboard/p/{ID}/new", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		selectedProject, err := service.Project(r.Context(), projectID)
		if err != nil {
			log.Println(err)
		}

		data := NewTaskPage{
			ID: projectID,
			DashboardPage: DashboardPage{
				Projects:        projects,
				SelectedProject: selectedProject,
			},
		}

		if err := RenderNewTaskPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
	})

	// SUBMIT NEW TASK
	r.Post("/dashboard/p/{ID}/new", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		// todo: proper error handling
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}

		// todo: Sanitize user input
		formTaskTitle := r.Form.Get("task-title")

		newTask, err := NewTaskWithID()
		if err != nil {
			log.Println(err)
		}

		// todo: grab *all* input from form
		newTask.Title = formTaskTitle
		newTask.Status = StatusDoing.String()
		newTask.Priority = DegreeHigh.String()
		newTask.Importance = DegreeHigh.String()
		newTask.DateCreated = time.Now()

		err = service.CreateTask(r.Context(), projectID, newTask)
		if err != nil {
			log.Println(err)
		}

		// failure path:
		// todo: return to same page with error message

		// ...

		// success path:
		http.Redirect(w, r, fmt.Sprintf("/dashboard/p/%s", projectID), http.StatusSeeOther)
	})

	//r.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(StaticFS))).ServeHTTP)
	r.HandleFunc("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("."))).ServeHTTP)
}
