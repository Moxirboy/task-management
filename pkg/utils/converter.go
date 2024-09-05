package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func ToInt(data interface{}) (n int, err error) {
	switch t := data.(type) {
	case int:
		return t, nil
	case string:
		return strconv.Atoi(t)
	default:
		return 0,
			errors.New(fmt.Sprintf("Invalid type received: %T", t))
	}
}

func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
