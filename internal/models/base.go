package models

import "time"

type At struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type By struct {
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}
