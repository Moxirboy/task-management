package jwt

import "github.com/gbrlsnchs/jwt/v3"

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	jwt.Payload
	Role string `json:"role"`
}
