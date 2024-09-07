package postgres

const (
	createUser = `
		INSERT INTO users
		(first_name,
		 last_name,
		 email,
		 position,
		 password
		 )
		VALUES ($1, $2, $3, $4, $5)
		returning id
`

	GetUserByAuthCred = `
		SELECT  id,
				first_name,
				last_name,
				position
		FROM users
		WHERE deleted_at IS NULL
		AND email = $1
		AND password = $2
	`
	checkFieldEmployee = `
		SELECT count(1)
		FROM users
		WHERE %s = $1 AND deleted_at IS NULL
	`
)
