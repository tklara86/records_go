package models

import (
	"database/sql"
	"time"
)

type Artist struct {
	ArtistID  int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordArtist struct {
	ID        int64
	ArtistID  int64
	RecordID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArtistModel struct {
	DB *sql.DB
}

func (m *ArtistModel) GetAll() ([]*Artist, error) {
	stmt := `
		SELECT * FROM artists
	`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	artists := []*Artist{}

	for rows.Next() {
		a := &Artist{}

		err := rows.Scan(&a.ArtistID, &a.Name, &a.CreatedAt, &a.UpdatedAt)
		if err != nil {
			return nil, err
		}

		artists = append(artists, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}

func (m *ArtistModel) Insert(a *Artist) (int, error) {
	stmt := `
		INSERT INTO artists (name, created_at, updated_at)
			VALUES (?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
	`

	result, err := m.DB.Exec(stmt, a.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *ArtistModel) InsertRecordArtist(ra []RecordArtist) (int, error) {
	stmt := `
		INSERT INTO record_artists (artist_id, record_id, created_at, updated_at)
		VALUES `

	args := []interface{}{}

	for _, v := range ra {
		args = append(args, v.ArtistID, v.RecordID)

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

func (m *ArtistModel) GetRecordArtist(recordID int) ([]*Artist, error) {
	stmt := `
		SELECT name FROM artists a
		LEFT JOIN record_artists ra ON ra.artist_id = a.artist_id
		LEFT JOIN records r ON r.id = ra.record_id
		WHERE r.id = ?;
	`

	rows, err := m.DB.Query(stmt, recordID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recordArtists := []*Artist{}

	for rows.Next() {
		a := &Artist{}

		err := rows.Scan(&a.Name)
		if err != nil {
			return nil, err
		}
		recordArtists = append(recordArtists, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recordArtists, nil

}
