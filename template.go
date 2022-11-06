package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
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

// DESIGN

func RenderLoginDesignPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_page_design_login.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}

func Render404DesignPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_page_design_404.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}

func RenderLandingDesignPage(w http.ResponseWriter) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_page_design_landing.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, nil)
	})
}

// APP

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
	BaseURL     string
	Redirect    string
	QueryParams string
}

func appendNonEmptyQueryParam(s []string, name, value string) {
	if len(strings.TrimSpace(name)) != 0 && len(strings.TrimSpace(value)) != 0 {
		s = append(s, fmt.Sprintf("%s=%s", name, value))
	}
}

func BuildQueryParams(redirect string) string {

	var params []string
	//for k, v := range vv {
	appendNonEmptyQueryParam(params, "redirect", redirect)
	//}

	if len(params) != 0 {
		return fmt.Sprint("?", strings.Join(params, "&"))
	}

	return ""
}

type ListTaskPage struct {
	DashboardPage
	ID    string
	Title string
	Tasks []Task
}

func RenderListTaskPage(w http.ResponseWriter, data ListTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_task_list_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type NewTaskPage struct {
	DashboardPage
	ID string
}

func RenderNewTaskPage(w http.ResponseWriter, data NewTaskPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_task_new_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}

type DashboardPage struct {
	Page
	Projects        []Project
	SelectedProject Project
}

func RenderDashboardPage(w http.ResponseWriter, data DashboardPage) error {
	return RenderHTML(w, func(wr io.Writer) error {
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_blank_page.html")
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
		tmpl, err := template.ParseFiles("tmpl_layout_dashboard.html", "tmpl_project_new_page.html")
		if err != nil {
			return err
		}
		return tmpl.Execute(w, &data)
	})
}
