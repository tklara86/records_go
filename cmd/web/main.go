package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/tklara86/records_go/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	records       *models.RecordModel
	genres        *models.GenreModel
	artists       *models.ArtistModel
	labels        *models.LabelModel
	tracklists    *models.TracklistModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":8000", "HTTP Netword Address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "Error\t", log.LstdFlags)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		errorLog.Println(err)
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// init a new template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		records:       &models.RecordModel{DB: db},
		genres:        &models.GenreModel{DB: db},
		artists:       &models.ArtistModel{DB: db},
		labels:        &models.LabelModel{DB: db},
		tracklists:    &models.TracklistModel{DB: db},
		templateCache: templateCache,
	}

	srv := http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server starting on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
