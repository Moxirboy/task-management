package usecase

import (
	"context"
	"food-delivery/internal/configs"
	"food-delivery/internal/models"
	"food-delivery/internal/service/storage/repo"
	"food-delivery/pkg/jwt"
	"food-delivery/pkg/logger"
	"strings"
	"time"
)

type authUseCase struct {
	authRepo repo.IAuthTokenRepository
	log      logger.Logger
}

func NewAuthUseCase(
	authRepo repo.IAuthTokenRepository,
	log logger.Logger,
) IAuthUseCase {
	return &authUseCase{
		authRepo: authRepo,
		log:      log,
	}
}

func (uc authUseCase) New(
	ctx context.Context,
	id, role string,
) (tokens *models.Tokens, err error) {
	return uc.generate(ctx, id, role)
}

func (uc authUseCase) Check(
	_ context.Context,
	accessToken string,
) (id, role string, err error) {
	data, err := jwt.ExtractTokenMetadata(accessToken)
	if err != nil {
		uc.log.Error("Error while checking access token:", err.Error())
		return
	}

	return data.JWTID, data.Role, nil
}

func (uc authUseCase) ReNew(
	ctx context.Context,
	refreshToken string,
) (tokens *models.Tokens, err error) {
	if err = uc.authRepo.CleanUp(ctx); err != nil {
		uc.log.Error("Error while cleaning up token in ReNew:", err.Error())
		return
	}

	at, err := uc.authRepo.Get(ctx, refreshToken)
	if err != nil {
		uc.log.Error(
			"Error while getting new token after renewing:",
			err.Error(),
		)
		return
	}

	err = uc.authRepo.Delete(ctx, refreshToken)
	if err != nil {
		uc.log.Error(
			"Error while deleting refresh token in ReNew:",
			err.Error(),
		)
		return
	}

	return uc.generate(ctx, at.ID, at.Role)
}

func (uc authUseCase) generate(
	ctx context.Context,
	id, role string,
) (tokens *models.Tokens, err error) {
	conf := configs.Load()

	tokens, err = jwt.GenerateNewTokens(id, role)
	if err != nil {
		uc.log.Error(
			"Error while generating new tokens in generate:",
			err.Error(),
		)
		return
	}
	now := time.Now()
	expireHours := conf.JWT.RefreshKeyExpireHours
	if strings.EqualFold(role, jwt.USER) {
		expireHours = conf.JWT.ClientRefreshExpireHours
	}

	err = uc.authRepo.Create(
		ctx,
		tokens.Refresh,
		id,
		role,
		now.Add(time.Hour*time.Duration(expireHours)),
	)
	if err != nil {
		uc.log.Error(
			"Error while creating token with its features:",
			err.Error(),
		)
		return
	}

	return tokens, nil
}
