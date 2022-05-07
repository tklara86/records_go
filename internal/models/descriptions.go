package models

import "time"

type Description struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordDescription struct {
	ID            int64
	DescriptionID int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
