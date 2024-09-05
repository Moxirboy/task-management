package postgres

import (
	"database/sql"
	"fmt"
	"sync"

	"food-delivery/internal/configs"

	_ "github.com/lib/pq" // pq for connection
	"github.com/pkg/errors"
)

var (
	instance *sql.DB
	once     sync.Once
)

// DB return database connection
func DB(cfg *configs.Postgres) (*sql.DB, error) {
	var err error
	once.Do(func() {
		psqlString := fmt.Sprintf(
			`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Database,
		)
		instance, err = sql.Open("postgres", psqlString)

	})

	if err != nil {
		return nil, errors.Wrap(err, "pgx.Connect")
	}

	err = instance.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "pg.Ping")
	}
	//forceMigration(instance, 1)

	return instance, nil
}
