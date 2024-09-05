package jwt

import (
	"food-delivery/internal/configs"
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(token string) (*TokenMetadata, error) {
	var (
		claims          TokenMetadata
		conf            = configs.Load()
		expValidator    = jwt.ExpirationTimeValidator(time.Now())
		validatePayload = jwt.ValidatePayload(&claims.Payload, expValidator)
	)
	_, err := jwt.Verify(
		[]byte(token),
		jwt.NewHS256([]byte(conf.JWT.SecretKey)),
		&claims,
		jwt.ValidateHeader,
		validatePayload,
	)

	if err != nil {
		return nil, err
	}

	return &claims, nil
}
