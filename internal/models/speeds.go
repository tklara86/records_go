package models

import "time"

type Speed struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordSpeed struct {
	ID        int64
	SpeedId   int64
	RecordID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
