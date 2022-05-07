package models

import "time"

type Size struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordSize struct {
	ID        int64
	SizeID    int64
	RecordID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
