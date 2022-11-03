package main

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

func RenderInternalServerError(w http.ResponseWriter, err error) {
	log.Printf("error: %+v\n", err)
}

func RenderHTML(w http.ResponseWriter, fn func(wr io.Writer) error) error {

	buf := new(bytes.Buffer)
	defer buf.Reset()

	if err := fn(buf); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write(buf.Bytes())
	return err
}

func RenderLandingPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_landing_page.html")
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
	Page
	ID    string
	Tasks []Task
}

func RenderListTaskPage(w http.ResponseWriter, data ListTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_app.html", "tmpl_task_list_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NewTaskPage struct {
	Page
	ID string
}

func RenderNewTaskPage(w http.ResponseWriter, data NewTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_app.html", "tmpl_task_new_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type ListProjectPage struct {
	Page
	Projects []Project
}

func RenderListProjectPage(w http.ResponseWriter, data ListProjectPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_project_list_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NewProjectPage struct {
	Page
}

func RenderNewProjectPage(w http.ResponseWriter, data NewProjectPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_project_new_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}
