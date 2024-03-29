package models

import (
	"database/sql"
	"errors"
	"time"
)

// Models
type Record struct {
	ID           int64
	Title        string
	ReleaseDate  string
	Image        string
	Status       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RecordArtist []*Artist
	RecordLabel  []*Label
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

	err := m.DB.QueryRow("SELECT * FROM records WHERE id = ?", recordId).Scan(&r.ID, &r.Title, &r.ReleaseDate, &r.Image, &r.Status, &r.CreatedAt, &r.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return r, nil

}

// SELECT name FROM artists a
// LEFT JOIN record_artists ra ON ra.artist_id = a.artist_id
// LEFT JOIN records r ON r.id = ra.record_id
// WHERE r.id = ?;

func (m *RecordModel) GetAll() ([]*Record, error) {
	stmt := `
		SELECT * FROM records
	`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// initialize empty slice to hold record structs
	records := []*Record{}

	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		r := &Record{}

		err := rows.Scan(&r.ID, &r.Title, &r.ReleaseDate, &r.Image, &r.Status, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Append it to the slice of records.
		records = append(records, r)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil

}
