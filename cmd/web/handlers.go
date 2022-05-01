package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// homeAdmin
func homeAdmin(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/common/navigation.tmpl",
		"./ui/html/pages/common/sidebar.tmpl",
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
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func createRecord(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create Record"))
}

func viewRecord(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display record with id of %d", id)
}

func viewRecords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All records"))
}
