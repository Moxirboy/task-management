package models

const (
	RoleClient = "CLIENT"
)

type Tokens struct {
	Access  string
	Refresh string
}
