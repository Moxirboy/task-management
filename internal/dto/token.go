package dto

import "time"

type AuthToken struct {
	Token    string     `json:"token"    db:"token"`
	ID       string     `json:"id"       db:"id"`
	Role     string     `json:"role"     db:"role"`
	Datetime *time.Time `json:"datetime" db:"datetime"`
}
