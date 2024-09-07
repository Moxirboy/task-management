package postgres

import (
	"context"
	"database/sql"
	"food-delivery/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

type authTokenRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewAuthTokenRepository(
	db *sql.DB,
	log logger.Logger,
) repo.IAuthTokenRepository {
	return &authTokenRepository{db: db, log: log}
}

func (r authTokenRepository) Create(
	ctx context.Context,
	token, id, role string,
	date time.Time,
) error {
	_, err := r.db.ExecContext(
		ctx,
		authCreate,
		token,
		date.Format(time.RFC3339),
		id,
		role,
	)
	if err != nil {
		r.log.Error(
			"Error while inserting into auth_token (Create):",
			err.Error(),
		)
		return err
	}

	return nil
}

func (r authTokenRepository) CleanUp(ctx context.Context) error {
	_, err := r.db.ExecContext(
		ctx,
		authCleanUp,
	)
	if err != nil {
		r.log.Error(
			"Error while deleting from auth_token (CleanUp):",
			err.Error(),
		)
		return errors.Wrap(err, "CleanUp")
	}

	return nil
}

func (r authTokenRepository) Get(
	ctx context.Context,
	token string,
) (*dto.AuthToken, error) {
	at := dto.AuthToken{}
	err := r.db.QueryRowContext(ctx, authGet, token, time.Now().Format(time.RFC3339)).
		Scan(
			&at.ID,
			&at.Token,
			&at.Role,
			&at.Datetime,
		)

	if err != nil {
		r.log.Error("Error while selecting from auth_token (Get):", err.Error())
		return nil, err
	}
	return &at, nil
}

func (r authTokenRepository) Delete(ctx context.Context, token string) error {
	_, err := r.db.ExecContext(
		ctx,
		authDelete,
		token,
	)
	if err != nil {
		r.log.Error(
			"Error while deleting from auth_token (Delete):",
			err.Error(),
		)
		return err
	}

	return nil
}
