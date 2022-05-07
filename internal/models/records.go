package models

import (
	"database/sql"
	"errors"
	"time"
)

// Models
type Record struct {
	ID          int64
	Title       string
	ReleaseDate string
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RecorImage struct {
	ID        int64
	RecordID  int64
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordModel struct {
	DB *sql.DB
}

func (m *RecordModel) Insert(rd *Record) (int, error) {
	stmt := `
		INSERT INTO records(title,release_date,image,created_at,updated_at) VALUES(?, ?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
	`
	result, err := m.DB.Exec(stmt, rd.Title, rd.ReleaseDate, rd.Image)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *RecordModel) Get(recordId int) (*Record, error) {
	r := &Record{}

	err := m.DB.QueryRow("SELECT * FROM records WHERE id = ?", recordId).Scan(&r.ID, &r.Title, &r.ReleaseDate, &r.Image, &r.CreatedAt, &r.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return r, nil

}
