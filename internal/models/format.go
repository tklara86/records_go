package models

import "time"

type Format struct {
	ID        int64
	Name      string
	Quantity  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordFormat struct {
	ID        int64
	FormatID  int64
	RecordID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
