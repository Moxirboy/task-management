package models

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Role string

const (
	UNAUTHORIZED Role = "unauthorized"
	BASIC        Role = "basic"
)

type AccessToken struct {
	ID   string
	Role string
}

func (r *AccessToken) MarshalBinary() ([]byte, error) {
	if r == nil {
		return nil, errors.New("nil pointer")
	}

	token := fmt.Sprintf("%s\t%s", r.ID, r.Role)

	return []byte(token), nil
}

func (r *AccessToken) UnmarshalBinary(data []byte) error {
	token := strings.Split(string(data), "\t")

	if len(token) != 2 {
		return errors.New("not supported format")
	}

	r.ID = token[0]
	r.Role = token[1]
	return nil
}
