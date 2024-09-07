package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"task-management/internal/models"
	"task-management/internal/service/storage/repo"
	"task-management/pkg/logger"
	"task-management/pkg/utils"
)

type userRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewUserRepository(db *sql.DB, log logger.Logger) repo.IUserRepository {
	return &userRepository{db: db, log: log}
}

func (u *userRepository) CreateUser(ctx context.Context, user *models.User) (string, error) {
	var (
		UserID string
	)
	row := u.db.QueryRowContext(
		ctx,
		createUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Position,
		user.Password,
	)
	if err := row.Scan(&UserID); err != nil {
		u.log.Error(
			"Error while inserting into employee_details",
			err.Error(),
		)
		return "", err
	}
	return UserID, nil
}

func (u *userRepository) Login(ctx context.Context, login, password string) (*models.User, error) {
	var (
		ID        string
		position  string
		firstname string
		lastname  sql.NullString
	)
	row := u.db.QueryRowContext(
		ctx,
		GetUserByAuthCred,
		login,
		password,
	)
	if err := row.Scan(&ID, &firstname, &lastname, &position); err != nil {
		u.log.Error(
			"Error while getting user",
			err.Error(),
		)
		return nil, err
	}
	return &models.User{
		ID:        ID,
		FirstName: firstname,
		LastName:  lastname.String,
		Position:  position,
		Email:     login,
	}, nil
}

func (u *userRepository) CheckField(
	ctx context.Context,
	field, value string,
) (bool, error) {
	var existsClient int

	if field == "position" {
		query := fmt.Sprintf(checkFieldEmployee, "position")
		row := u.db.QueryRowContext(ctx, query, value)
		if err := row.Scan(&existsClient); err != nil {
			u.log.Error(
				"Error while checking field 'position' of user",
				err.Error(),
			)
			return false, err
		}
	} else {
		return false, utils.ErrInvalidField
	}

	if existsClient > 0 {
		return false, nil
	}

	return true, nil
}
