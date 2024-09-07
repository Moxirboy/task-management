package usecase

import (
	"context"
	"food-delivery/internal/configs"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/logger"
	"food-delivery/pkg/utils"
	"strings"
)

type accountUseCase struct {
	UserRepo repo.IUserRepository
	log      logger.Logger
	cfg      *configs.Config
}

func NewAccountUseCase(
	UserRepo repo.IUserRepository,
	log logger.Logger,
	cfg *configs.Config,
) IAccountUseCase {
	return &accountUseCase{
		UserRepo: UserRepo,
		log:      log,
		cfg:      cfg,
	}
}

func (e *accountUseCase) CreateUser(
	ctx context.Context,
	User *models.User,
) (err error) {

	positions := strings.Split(User.Position, ",")

	for _, position := range positions {
		if position == string(models.PositionAdmin) {
			User.Password = utils.Hash([]byte(e.cfg.Setup.AdminPassword))
			isSuperAdminExist, errCheck := e.UserRepo.CheckField(
				ctx,
				"position",
				User.Position,
			)
			if errCheck != nil {
				e.log.Error(
					"Error while checking field of user position:",
					err.Error(),
				)
				return errCheck
			}

			if !isSuperAdminExist {
				return utils.ErrAlreadyExist
			}
		}

	}
	User.Password = utils.Hash([]byte(User.Password))
	User.ID, err = e.UserRepo.CreateUser(ctx, User)
	if err != nil {
		e.log.Error("Error while creating employee:", err.Error())
		return err
	}

	return nil
}

func (e *accountUseCase) LoginUser(
	ctx context.Context,
	login, password string,
) (res *models.User, err error) {
	password = utils.Hash([]byte(password))
	res, err = e.UserRepo.Login(ctx, login, password)
	if err != nil {
		e.log.Error("Error while checking login credentials:", err.Error())
		return nil, err
	}

	return res, nil
}
