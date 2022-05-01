package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/record/view", viewRecord)
	mux.HandleFunc("/record/create", createRecord)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalln(err)
	}

}
