package main

import (
	"html/template"
	"net/http"
)

// homeAdmin
func (app *application) homeAdmin(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
		"./ui/html/pages/admin/about.tmpl",
		"./ui/html/pages/admin/home.tmpl",
	}

	type Link struct {
		LinkTitle string
		LinkPath  string
	}

	type Links struct {
		Link []Link
	}

	data := Links{
		Link: []Link{
			{
				LinkTitle: "Home",
				LinkPath:  "/admin",
			},
			{
				LinkTitle: "Categories",
				LinkPath:  "/admin/categories/",
			},
		},
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) createRecordPost(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create Record"))
}

func (app *application) createRecordGet(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
		"./ui/html/pages/record/create.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) viewRecord(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
		"./ui/html/pages/admin/about.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, err)
	}

	// id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// if err != nil || id < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	//fmt.Fprintf(w, "Display record with id of %d", id)
}

func (app *application) viewRecords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All records"))
}
