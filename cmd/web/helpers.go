package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	// Initialize a new buffer.
	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)

}

type Link struct {
	LinkTitle string
	LinkPath  string
}

type SideBarLink struct {
	LinkTitle string
	LinkIcon  string
	LinkHref  string
}

type FormOption struct {
	FormName string
	FormType string
	FormFor  string
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		Links: []Link{
			{
				LinkTitle: "Home",
				LinkPath:  "/admin",
			},
			{
				LinkTitle: "Categories",
				LinkPath:  "/admin/categories/",
			},
		},
		SideBarLinks: []SideBarLink{
			{
				LinkTitle: "Home",
				LinkIcon:  "home",
				LinkHref:  "/admin",
			},
			{
				LinkTitle: "Albums",
				LinkIcon:  "archive",
				LinkHref:  "/admin/records",
			},
		},
		FormOptions: []FormOption{
			{
				FormName: "Title",
				FormType: "text",
				FormFor:  "recordTitle",
			},
			{
				FormName: "Release Date",
				FormType: "text",
				FormFor:  "recordReleaseDate",
			},
			// {
			// 	FormName: "Label",
			// 	FormType: "text",
			// 	FormFor:  "recordLabel",
			// },
		},
	}
}
