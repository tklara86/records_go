package models

import "time"

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
