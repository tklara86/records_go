package models

import "time"

type Label struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LabelCatalogueNumber struct {
	ID              int64
	LabelID         int64
	CatalogueNumebr string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type LabelCatalogueNumberToRecord struct {
	ID                int64
	CatalogueNumberID int64
	RecordID          int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RecordLabel struct {
	ID        int64
	LabelID   int64
	RecordID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
