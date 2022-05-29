package models

import (
	"database/sql"
	"time"
)

type Label struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	InputName string    `json:"input_name"`
	CreatedAt time.Time `json:"-"` // ommit
	UpdatedAt time.Time `json:"-"`
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

type LabelModel struct {
	DB *sql.DB
}

func (m *LabelModel) GetAll() ([]*Label, error) {
	stmt := `
		SELECT * FROM labels
	`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	labels := []*Label{}

	for rows.Next() {
		l := &Label{}

		err := rows.Scan(&l.ID, &l.Name, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			return nil, err
		}

		labels = append(labels, l)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return labels, nil
}

func (m *LabelModel) Insert(l *Label) (int, error) {
	stmt := `
		INSERT INTO labels (name, created_at, updated_at)
			VALUES (?, UTC_TIMESTAMP(), UTC_TIMESTAMP())
	`

	result, err := m.DB.Exec(stmt, l.Name)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *LabelModel) InsertRecordLabel(rl []RecordLabel) (int, error) {
	stmt := `
		INSERT INTO record_labels (label_id, record_id, created_at, updated_at)
		VALUES `

	args := []interface{}{}

	for _, v := range rl {
		args = append(args, v.LabelID, v.RecordID)

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

func (m *LabelModel) GetLabelArtist(recordID int) ([]*Label, error) {
	stmt := `
		SELECT name FROM labels l
		LEFT JOIN record_labels rl ON rl.label_id = l.id
		LEFT JOIN records r ON r.id = rl.record_id
		WHERE r.id = ?;
	`

	rows, err := m.DB.Query(stmt, recordID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recordLabels := []*Label{}

	for rows.Next() {
		l := &Label{}

		err := rows.Scan(&l.Name)
		if err != nil {
			return nil, err
		}
		recordLabels = append(recordLabels, l)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return recordLabels, nil

}
