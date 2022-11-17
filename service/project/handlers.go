package project

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func NewProjectHandler(w http.ResponseWriter, r *http.Request, service Service) {
	redirect := r.URL.Query().Get("redirect")

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	data := NewProjectPage{
		DashboardPage: DashboardPage{
			Projects: projects,
			Redirect: redirect,
		},
	}

	if err := RenderNewProjectPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func SubmitNewProjectHandler(w http.ResponseWriter, r *http.Request, service Service) {
	// todo: proper error handling
	// todo: redirect to internal server error page?
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	// todo: should sanitzation of data be a domain concern?
	formProjectTitle := strings.TrimSpace(r.Form.Get("project-title"))

	newProject, _ := NewProject()
	newProject.Title = formProjectTitle

	// todo: proper error handling
	id, err := service.CreateProject(r.Context(), newProject)
	if err != nil {
		log.Println(err)
	}

	// failure path:
	// ...
	// todo: return to same page with error message if an error occurred
	// todo: replace with new project error template or add error detail struct to template and conditionally check

	// success path:
	http.Redirect(w, r, fmt.Sprintf("/dashboard/project/%s", id), http.StatusSeeOther)
}

func EditProjectHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	data := EditProjectPage{
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
			Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
		},
		TargetProject: currentProject,
	}

	if err := RenderEditProjectPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func SubmitUpdatedProjectHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	// todo: proper error handling
	// todo: redirect to internal server error page?
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	// todo: should sanitzation of data be a domain concern?
	formProjectTitle := strings.TrimSpace(r.Form.Get("project-title"))
	//formProjectDescription := strings.TrimSpace(r.Form.Get("project-description"))

	var formErrors []string

	titleValidErr := TitleValid(formProjectTitle)

	switch titleValidErr {
	case ErrTitleEmpty:
		formErrors = append(formErrors, "ReadProject title is empty. Please provide a title.")
		break
	case ErrTitleOverMaxLength:
		formErrors = append(formErrors, "ReadProject title has reached max length. Please provide a title under 201 characters long.")
		break
	}

	// failure path:
	if len(formErrors) > 0 {

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		data := EditProjectPage{
			DashboardPage: DashboardPage{
				Projects:       projects,
				CurrentProject: currentProject,
				Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
			},
			TargetProject: Project{
				ID:    currentProject.ID,
				Title: formProjectTitle,
				//Description: formProjectDescription,
			},

			Errors: formErrors,
		}

		if err := RenderEditProjectPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	// success path:
	currentProject.Title = formProjectTitle
	//currentProject.Description = formProjectDescription

	err = service.UpdateProject(r.Context(), currentProject)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard/project/%s", projectID), http.StatusSeeOther)
}

func DeleteProjectConfirmationHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	data := DeleteProjectPage{
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
			Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
		},
	}

	if err := RenderDeleteProjectPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func DeleteProjectHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	err := service.DeleteProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		RenderInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func ProjectHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
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
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
			Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
		},
		Tasks: tasks,
	}
	if err := RenderListTaskPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func NewTaskHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	data := NewTaskPage{
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
		},
	}

	if err := RenderNewTaskPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func SubmitNewTaskHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string) {

	_, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	// todo: proper error handling
	// todo: redirect to internal server error page?
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	// todo: should sanitzation of data be a domain concern?
	formTaskTitle := strings.TrimSpace(r.Form.Get("task-title"))
	formTaskStatus := strings.TrimSpace(r.Form.Get("task-status"))
	formTaskPriority := strings.TrimSpace(r.Form.Get("task-priority"))
	formTaskImportance := strings.TrimSpace(r.Form.Get("task-importance"))

	newTask, err := NewTask()
	if err != nil {
		log.Println(err)
	}

	// todo: grab *all* input from form
	newTask.Title = formTaskTitle
	newTask.Status = formTaskStatus
	newTask.Priority = formTaskPriority
	newTask.Importance = formTaskImportance
	newTask.DateCreated = time.Now()

	_, err = service.CreateTask(r.Context(), projectID, newTask)
	if err != nil {
		log.Println(err)
	}

	// failure path:
	// todo: return to same page with error message

	// ...

	// success path:
	http.Redirect(w, r, fmt.Sprintf("/dashboard/project/%s", projectID), http.StatusSeeOther)
}

func DeleteTaskConfirmationHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string, taskID string) {
	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	task, err := service.ReadTask(r.Context(), projectID, taskID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	data := DeleteTaskPage{
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
			Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
		},
		Task: task,
	}

	if err := RenderDeleteTaskPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string, taskID string) {
	err := service.DeleteTask(r.Context(), projectID, taskID)
	if err != nil {
		log.Println(err)
		RenderInternalServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard/project/%s", projectID), http.StatusSeeOther)
}

func EditTaskHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string, taskID string) {

	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	projects, err := service.ListProjects(r.Context())
	if err != nil {
		log.Println(err)
	}

	task, err := service.ReadTask(r.Context(), projectID, taskID)
	if err != nil {
		log.Println(err)
	}

	data := EditTaskPage{
		DashboardPage: DashboardPage{
			Projects:       projects,
			CurrentProject: currentProject,
			Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
		},
		Task:       task,
		TargetTask: task,
	}

	if err := RenderEditTaskPage(w, data); err != nil {
		RenderInternalServerError(w, err)
	}
}

func SubmitUpdatedTaskHandler(w http.ResponseWriter, r *http.Request, service Service, projectID string, taskID string) {

	currentProject, err := service.ReadProject(r.Context(), projectID)
	if err != nil {
		log.Println(err)
		if err := RenderProject404Page(w, NotFoundPage{}); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	// todo: proper error handling
	// todo: redirect to internal server error page?
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}

	// todo: should sanitzation of data be a domain concern?
	formTaskTitle := strings.TrimSpace(r.Form.Get("task-title"))
	formTaskStatus := strings.TrimSpace(r.Form.Get("task-status"))
	formTaskPriority := strings.TrimSpace(r.Form.Get("task-priority"))
	formTaskImportance := strings.TrimSpace(r.Form.Get("task-importance"))

	var formErrors []string

	titleValidErr := TitleValid(formTaskTitle)

	switch titleValidErr {
	case ErrTitleEmpty:
		formErrors = append(formErrors, "ReadTask title is empty. Please provide a title.")
		break
	case ErrTitleOverMaxLength:
		formErrors = append(formErrors, "ReadTask title has reached max length. Please provide a title under 201 characters long.")
		break
	}

	// failure path:
	if len(formErrors) > 0 {

		projects, err := service.ListProjects(r.Context())
		if err != nil {
			log.Println(err)
		}

		task, err := service.ReadTask(r.Context(), projectID, taskID)
		if err != nil {
			log.Println(err)
		}

		data := EditTaskPage{
			DashboardPage: DashboardPage{
				Projects:       projects,
				CurrentProject: currentProject,
				Redirect:       fmt.Sprintf("/dashboard/project/%s", projectID),
			},
			Task: task,
			TargetTask: Task{
				ID:          task.ID,
				Title:       formTaskTitle,
				Status:      formTaskStatus,
				Priority:    formTaskPriority,
				Importance:  formTaskImportance,
				DateCreated: task.DateCreated,
			},

			Errors: formErrors,
		}

		if err := RenderEditTaskPage(w, data); err != nil {
			RenderInternalServerError(w, err)
		}
		return
	}

	// success path:
	existingTask, err := service.ReadTask(r.Context(), projectID, taskID)
	if err != nil {
		log.Println(err)
	}

	existingTask.Title = formTaskTitle
	existingTask.Status = formTaskStatus
	existingTask.Priority = formTaskPriority
	existingTask.Importance = formTaskImportance

	err = service.UpdateTask(r.Context(), projectID, existingTask)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, fmt.Sprintf("/dashboard/project/%s", projectID), http.StatusSeeOther)
}
