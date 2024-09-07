package v1

import (
	"context"
	"fmt"
	"log"
	"task-management/internal/dto"
	"task-management/internal/models"
	"task-management/internal/service/usecase"
	"task-management/pkg/logger"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related operations.
type AuthHandler struct {
	log       logger.Logger
	AccountUC usecase.IAccountUseCase
	AuthUC    usecase.IAuthUseCase
}

// NewAuthHandler initializes the authentication routes.
//
//	@BasePath	/api/v1
func NewAuthHandler(
	r *gin.RouterGroup,
	l logger.Logger,
	AuthUC usecase.IAuthUseCase,
	AccountUC usecase.IAccountUseCase,
) {
	handler := &AuthHandler{
		log:       l,
		AuthUC:    AuthUC,
		AccountUC: AccountUC,
	}
	auths := r.Group("/auth")

	auths.POST("/login", handler.login)
	auths.POST("/refresh", handler.renew)
	auths.POST("/sign", handler.signUp)
	auths.GET("/set", handler.SetHandler)
}

// renew handles the token refresh operation.
//
//	@Summary		Refresh access token
//	@Description	Refreshes the access token using a refresh token.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.RenewRequest	true	"Refresh Token"
//	@Success		200		{object}	dto.BaseResponse	"Token refreshed successfully"
//	@Failure		400		{object}	dto.BaseResponse"Invalid request"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/auth/refresh [post]
func (a *AuthHandler) renew(c *gin.Context) {
	var body dto.RenewRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	res, err := a.AuthUC.ReNew(
		context.Background(),
		body.RefreshToken,
	)

	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(
		c,
		toRenewResponse(c.Request.Context(), res),
		nil,
	)
}

// login handles user login and token generation.
//
//	@Summary		User login
//	@Description	Authenticates a user and generates access and refresh tokens.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	dto.BaseResponse	"Login successful"
//	@Failure		400		{object}	dto.BaseResponse	"Invalid login request"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/auth/login [post]
func (h *AuthHandler) login(c *gin.Context) {
	var body dto.LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error parsing body: ", err)
		utils.SendResponse(c, nil, err)
		return
	}

	invalidParams := utils.Validate(body)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}

	user, err := h.AccountUC.LoginUser(
		c.Request.Context(),
		body.Login,
		body.Password,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	tokens, err := h.AuthUC.New(c.Request.Context(),
		user.ID,
		user.Position,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c,
		toLoginResponse(c.Request.Context(), user, tokens),
		nil,
	)
}

// signUp handles user registration.
//
//	@Summary		User sign up
//	@Description	Registers a new user and generates access and refresh tokens.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.SignUpRequest	true	"User sign-up details"
//	@Success		200		{object}	dto.BaseResponse	"Sign-up successful"
//	@Failure		400		{object}	dto.BaseResponse	"Invalid sign-up request"
//	@Failure		500		{object}	dto.BaseResponse	"Internal server error"
//	@Router			/auth/sign [post]
func (h *AuthHandler) signUp(c *gin.Context) {
	var body dto.SignUpRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error parsing body: ", err)
		utils.SendResponse(c, nil, err)
		return
	}

	invalidParams := utils.Validate(body)
	if invalidParams != nil {
		utils.SendResponse(c, invalidParams, nil)
		return
	}
	user := models.NewUser(body)
	err := h.AccountUC.CreateUser(
		c.Request.Context(),
		user,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	tokens, err := h.AuthUC.New(c.Request.Context(),
		user.ID,
		user.Position,
	)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c,
		toLoginResponse(c.Request.Context(), user, tokens),
		nil,
	)
}

// @Summary Set Authorization Header
// @Description Sets the Authorization header with a Bearer token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param token query string true "Authorization Token"
// @Success 200 {object} map[string]string "Success"
// @Router /auth/set [get]
func (h *AuthHandler) SetHandler(c *gin.Context) {
	query := c.Query("token")
	bear := fmt.Sprintf("Bearer %s", query)
	c.Header("Authorization", bear)
	c.JSON(200, gin.H{"message": "Authorization header set"})
}
