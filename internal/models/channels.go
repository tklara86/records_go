package models

import "time"

type Channel struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordChannel struct {
	ID        int64
	ChannelID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
