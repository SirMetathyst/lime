package main

import (
	"github.com/SirMetathyst/lime/service/project"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Route(r *chi.Mux, service project.Service) {

	// todo: turn the app handlers into http middleware or handlers

	// LANDING PAGE
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := project.RenderLandingPage(w); err != nil {
			project.RenderInternalServerError(w, err)
		}
	})

	// LOGIN PAGE
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := project.RenderLoginPage(w); err != nil {
			project.RenderInternalServerError(w, err)
		}
	})

	// TESTING: REDIRECT TO NEW PROJECT PAGE
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard/project/create", http.StatusSeeOther)
	})

	// TESTING: DASHBOARD: REDIRECT TO NEW PROJECT PAGE
	r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard/project/create", http.StatusSeeOther)
	})

	// TESTING: DASHBOARD: POST FOR LOGIN
	r.Post("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard/project/create", http.StatusSeeOther)
	})

	// DASHBOARD: NEW PROJECT PAGE
	r.Get("/dashboard/project/create", func(w http.ResponseWriter, r *http.Request) {

		project.NewProjectHandler(w, r, service)
	})

	// DASHBOARD: SUBMIT NEW PROJECT
	r.Post("/dashboard/project/create", func(w http.ResponseWriter, r *http.Request) {

		project.SubmitNewProjectHandler(w, r, service)
	})

	// DASHBOARD: EDIT PROJECT PAGE
	r.Get("/dashboard/project/{PROJECT_ID}/edit", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")

		project.EditProjectHandler(w, r, service, projectID)
	})

	// DASHBOARD: SUBMIT UPDATED PROJECT
	r.Post("/dashboard/project/{PROJECT_ID}/edit", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")

		project.SubmitUpdatedProjectHandler(w, r, service, projectID)
	})

	// DASHBOARD: DELETE PROJECT CONFIRMATION PAGE
	r.Get("/dashboard/project/{PROJECT_ID}/delete-confirm", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")

		project.DeleteProjectConfirmationHandler(w, r, service, projectID)
	})

	// DASHBOARD: DELETE PROJECT
	r.Post("/dashboard/project/{PROJECT_ID}/delete", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")

		project.DeleteProjectHandler(w, r, service, projectID)
	})

	// DASHBOARD: PROJECT PAGE
	r.Get("/dashboard/project/{ID}", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		project.ProjectHandler(w, r, service, projectID)
	})

	// DASHBOARD: NEW TASK PAGE
	r.Get("/dashboard/project/{ID}/task/create", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		project.NewTaskHandler(w, r, service, projectID)
	})

	// DASHBOARD: SUBMIT NEW TASK
	r.Post("/dashboard/project/{ID}/task/create", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "ID")

		project.SubmitNewTaskHandler(w, r, service, projectID)
	})

	// DASHBOARD: DELETE TASK CONFIRMATION PAGE
	r.Get("/dashboard/project/{PROJECT_ID}/{TASK_ID}/delete-confirm", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")
		taskID := chi.URLParam(r, "TASK_ID")

		project.DeleteTaskConfirmationHandler(w, r, service, projectID, taskID)
	})

	// DASHBOARD: DELETE TASK
	r.Post("/dashboard/project/{PROJECT_ID}/{TASK_ID}/delete", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")
		taskID := chi.URLParam(r, "TASK_ID")

		project.DeleteTaskHandler(w, r, service, projectID, taskID)
	})

	// DASHBOARD: EDIT TASK PAGE
	r.Get("/dashboard/project/{PROJECT_ID}/task/{TASK_ID}/edit", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")
		taskID := chi.URLParam(r, "TASK_ID")

		project.EditTaskHandler(w, r, service, projectID, taskID)
	})

	// DASHBOARD: SUBMIT UPDATED TASK
	r.Post("/dashboard/project/{PROJECT_ID}/task/{TASK_ID}/edit", func(w http.ResponseWriter, r *http.Request) {

		projectID := chi.URLParam(r, "PROJECT_ID")
		taskID := chi.URLParam(r, "TASK_ID")

		project.SubmitUpdatedTaskHandler(w, r, service, projectID, taskID)
	})

	//r.HandleFunc("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(StaticFS))).ServeHTTP)
	r.HandleFunc("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("."))).ServeHTTP)
}
