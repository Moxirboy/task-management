package utils

import (
	"database/sql"
	"errors"
	"net/http"

	"food-delivery/internal/dto"
	"github.com/gin-gonic/gin"
)

const (
	ErrMessageUnknown  = "Unknown server error"
	ErrMessageNotFound = "Not found"
)

const (
	StatusNotFound            = "NOT_FOUND"
	StatusInternalServerError = "INTERNAL_SERVER_ERROR"
	StatusInvalidField        = "INVALID_FIELD"
	StatusAlreadyExist        = "ALREADY_EXIST"
	NotAuthenticated          = "NOT_AUTHENTICATED"
)

var (
	ErrInvalidField = errors.New("invelid field")
	ErrAlreadyExist = errors.New("already exist")
	ErrNotAuthenticated = errors.New("not authenticated")
)

func SendResponse(
	ctx *gin.Context,
	response interface{},
	err error,
) {
	if err != nil {
		code, status, message := parseError(err)
		ctx.JSON(
			code,
			dto.BaseResponse{
				Status:  status,
				Message: message,
			},
		)
		return
	}

	switch response.(type) {
	case []*dto.InvalidParams:
		ctx.JSON(
			http.StatusBadRequest,
			dto.BaseResponse{
				Data: response,
			},
		)
	default:
		ctx.JSON(
			http.StatusOK,
			dto.BaseResponse{
				Data: response,
			},
		)
	}
}

func parseError(err error) (code int, status, message string) {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, StatusNotFound, ErrMessageNotFound
	} else if errors.Is(err, ErrInvalidField) {
		return http.StatusBadRequest, StatusInvalidField, ""
	} else if errors.Is(err, ErrAlreadyExist) {
		return http.StatusBadRequest, StatusAlreadyExist, ""
	}
	return http.StatusInternalServerError, StatusInternalServerError, err.Error()
}
