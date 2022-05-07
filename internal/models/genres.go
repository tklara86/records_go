package models

import (
	"database/sql"
	"time"
)

type Genre struct {
	GenreID   int64
	GenreName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordGenre struct {
	ID        int64
	RecordID  int64
	GenreID   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GenreModel struct {
	DB *sql.DB
}
