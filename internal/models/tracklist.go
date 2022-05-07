package models

import "time"

type Tracklist struct {
	ID            int64
	Track         string
	TrackTitle    string
	TrackDuration string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RecordTracklist struct {
	ID          int64
	RecordID    int64
	TracklistID int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RecordTracklistArtist struct {
	ID          int64
	TracklistID int64
	ArtistID    int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
