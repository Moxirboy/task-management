package postgres

const (
	authCreate = `
				INSERT INTO
					auth_token(token, datetime, id, role)
				VALUES($1, $2, $3, $4)`

	authCleanUp = `DELETE FROM auth_token WHERE datetime < CURRENT_TIMESTAMP`

	authGet = `
			SELECT
				id, token, role, datetime
			FROM
				auth_token
			WHERE
				token = $1 AND datetime::timestamp > $2::timestamp`

	authDelete = `DELETE FROM auth_token WHERE token = $1`
)
