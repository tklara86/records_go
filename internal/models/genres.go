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

func (m *GenreModel) Insert(g *Genre) (int, error) {
	stmt := `
		INSERT INTO genres(genre_name, created_at, updated_at)
		VALUES(?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
	`

	result, err := m.DB.Exec(stmt, g.GenreName)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (m *GenreModel) InsertRecordGenre(rg []RecordGenre) (int, error) {
	stmt := `
		INSERT INTO record_genres (record_id, genre_id, created_at, updated_at)
		VALUES `

	args := []interface{}{}

	for _, v := range rg {
		args = append(args, v.RecordID, v.GenreID)
		numFields := 1

		for j := 0; j < numFields; j++ {
			stmt += `(?,` + `?` + `, UTC_TIMESTAMP(), UTC_TIMESTAMP()),`
		}
		stmt = stmt[:len(stmt)-1] + `,`
	}
	stmt = stmt[:len(stmt)-1]

	result, err := m.DB.Exec(stmt, args...)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
