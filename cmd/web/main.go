package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8000", "HTTP Netword Address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "Error\t", log.LstdFlags)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server starting on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
