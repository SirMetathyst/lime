package project

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
)

///go:embed *.html
//var FS embed.FS

///go:embed *.css
//var StaticFS embed.FS

var (
//LandingPageTemplate = template.Must(template.ParseFS(FS, "tmpl_landing_page.html"))
)

// todo: change how templates are found. fix: hard coded directory

func RenderInternalServerError(w http.ResponseWriter, err error) {
	log.Printf("error: %+v\n", err)
}

type DashboardPage struct {
	Page
	Projects       []Project
	CurrentProject Project
	Redirect       string
}

func RenderHTML(w http.ResponseWriter, fn func(wr io.Writer) error) error {

	buf := new(bytes.Buffer)
	defer buf.Reset()

	if err := fn(buf); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")

	_, err := w.Write(buf.Bytes())
	return err
}

func RenderLoginPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_page_login.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}

func RenderLandingPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_page_landing.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}

type Page struct {
	BaseURL string
}

type ListTaskPage struct {
	DashboardPage
	Tasks []Task
}

// todo: remove
func RenderListTaskPage(w http.ResponseWriter, data ListTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_task_list.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type EditTaskPage struct {
	DashboardPage
	Task       Task
	TargetTask Task
	Errors     []string
}

// todo: remove
func RenderEditTaskPage(w http.ResponseWriter, data EditTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_task_edit.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type DeleteTaskPage struct {
	DashboardPage
	Task Task
}

// todo: remove
func RenderDeleteTaskPage(w http.ResponseWriter, data DeleteTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/template/tmpl_layout_dashboard.html", "service/project/tmpl_page_task_delete_confirm.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NewTaskPage struct {
	DashboardPage
	Errors []string
}

// todo: remove
func RenderNewTaskPage(w http.ResponseWriter, data NewTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_task_new.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NewProjectPage struct {
	DashboardPage
}

func RenderNewProjectPage(w http.ResponseWriter, data NewProjectPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_project_new.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type EditProjectPage struct {
	DashboardPage
	TargetProject Project
	Errors        []string
}

func RenderEditProjectPage(w http.ResponseWriter, data EditProjectPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_project_edit.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type DeleteProjectPage struct {
	DashboardPage
}

func RenderDeleteProjectPage(w http.ResponseWriter, data DeleteProjectPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_layout_dashboard.html", "service/project/tmpl_page_project_delete_confirm.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NotFoundPage struct {
	Page
}

func RenderProject404Page(w http.ResponseWriter, data NotFoundPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("service/project/tmpl_page_project_404.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}
