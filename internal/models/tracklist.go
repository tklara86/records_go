package models

import (
	"database/sql"
	"time"
)

type Tracklist struct {
	ID            int64
	RecordID      int64
	Track         string
	TrackTitle    string
	TrackDuration string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RecordTracklistArtist struct {
	ID          int64
	TracklistID int64
	ArtistID    int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TracklistModel struct {
	DB *sql.DB
}

// func (m *TracklistModel) Insert(t []Tracklist) (int, error) {
// 	stmt := `
// 		INSERT INTO tracklist (record_id, track, track_title, track_duration, created_at, updated_at)
// 		VALUES `

// 	args := []interface{}{}

// 	for _, k := range t {
// 		args = append(args, k.RecordID, k.Track, k.TrackTitle, k.TrackDuration)
// 		numFields := 1

// 		for j := 0; j < numFields; j++ {
// 			stmt += `(?,` + `?,` + `?,` + `?` + `, UTC_TIMESTAMP(), UTC_TIMESTAMP()),`
// 		}
// 		stmt = stmt[:len(stmt)-1] + `,`
// 	}
// 	stmt = stmt[:len(stmt)-1]

// 	result, err := m.DB.Exec(stmt, args...)
// 	if err != nil {
// 		return 0, nil
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, nil
// 	}

// 	return int(id), nil
// }

func (m *TracklistModel) Insert(t []Tracklist) ([]int64, error) {
	stmt := `
		INSERT INTO tracklist (record_id, track, track_title, track_duration, created_at, updated_at)
		VALUES (?,?,?,?,UTC_TIMESTAMP(), UTC_TIMESTAMP())`

	var insertedIds []int64

	for _, k := range t {
		result, err := m.DB.Exec(stmt, k.RecordID, k.Track, k.TrackTitle, k.TrackDuration)
		if err != nil {
			return nil, nil
		}

		id, err := result.LastInsertId()
		if err != nil {
			return nil, nil
		}

		insertedIds = append(insertedIds, id)

	}

	return insertedIds, nil

}

func (m *TracklistModel) InsertRecordTracklistArtist(rta []RecordTracklistArtist) (int, error) {
	stmt := `INSERT INTO record_tracklists_artists (tracklist_id, artist_id, created_at, updated_at)
	VALUES `

	args := []interface{}{}

	for _, k := range rta {
		args = append(args, k.TracklistID, k.ArtistID)
		numFields := 1

		for j := 0; j < numFields; j++ {
			stmt += `(?,` + `?` + `, UTC_TIMESTAMP(), UTC_TIMESTAMP()),`
		}
		stmt = stmt[:len(stmt)-1] + `,`
	}
	stmt = stmt[:len(stmt)-1]

	result, err := m.DB.Exec(stmt, args...)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return int(id), nil
}
