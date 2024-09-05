package jwt

import (
	"encoding/base32"

	"crypto/rand"
	"fmt"
	"food-delivery/internal/configs"
	"food-delivery/internal/models"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var (
	ErrInvalidUUID = errors.New("invalid UUID format")
)

const (
	USER = "USER"
)

// GenerateNewTokens func for generate a new Access & Refresh tokens.
func GenerateNewTokens(id, role string) (*models.Tokens, error) {
	// Generate Parse Access token.
	accessToken, err := generateNewAccessToken(id, role)
	if err != nil {
		// Return token generation error.
		return nil, errors.Wrap(err, "generate access token")
	}

	// Generate Parse Refresh token.
	refreshToken, err := generateNewRefreshToken(id, role)
	if err != nil {
		// Return token generation error.
		return nil, errors.Wrap(err, "generate refresh token")
	}

	return &models.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id, role string) (string, error) {
	conf := configs.Load()

	if !isValidUUID(id) {
		return "", ErrInvalidUUID
	}

	now := time.Now()

	expiresAt := now.
		Add(time.Minute * time.Duration(conf.JWT.SecretKeyExpireMinutes))

	// Create a new claims.
	claims := TokenMetadata{
		Role: role,
		Payload: jwt.Payload{
			JWTID:          id,
			ExpirationTime: jwt.NumericDate(expiresAt),
			NotBefore:      jwt.NumericDate(time.Now()),
		},
	}

	// Create a new Parse access token with claims.
	token, err := jwt.Sign(
		claims,
		jwt.NewHS256([]byte(configs.Load().JWT.SecretKey)),
	)
	if err != nil {
		// Return error, it Parse token generation failed.
		return "", err
	}

	return string(token), nil
}

func generateNewRefreshToken(id, role string) (string, error) {
	conf := configs.Load()
	now := time.Now()
	expireHours := conf.JWT.RefreshKeyExpireHours
	if strings.EqualFold(role, USER) {
		expireHours = conf.JWT.ClientRefreshExpireHours
	}
	token, err := jwt.Sign(
		jwt.Payload{
			JWTID:     id,
			NotBefore: jwt.NumericDate(now),
			ExpirationTime: jwt.NumericDate(now.
				Add(time.Hour * time.Duration(expireHours))),
		},
		jwt.NewHS256([]byte(configs.Load().JWT.RefreshKey)),
	)
	if err != nil {
		return "", err
	}

	res := fmt.Sprintf(
		"%s.%s.%d",
		string(token),
		getToken(12),
		now.Unix(),
	)

	return res, nil
}

func getToken(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
