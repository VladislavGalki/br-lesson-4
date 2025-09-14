package auth

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

/*
<header>.<payload>.<signature>

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJzdWIiOiIxMjMiLCJpc3MiOiJteS1hcGkiLCJhdWQiOiJteS.
hNrgH0pB4B1...  (подпись)

header - {alg: "HS256", typ: "JWT", kid: "key-1"}
payload - {sub: "123", iss: "my-api", aud: "my"}
Стандартные поля:
iss (issuer) - кто выдал токен
aud (audience) - кому предназначен
sub (subject) - кто это
exp - когда истекает время жизни токена ( unix сек)
nbf - не раньше чем
iat - когда был создан
jti - уникальный идентификатор

Signing Method
HMAC using S
*/

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type HS256Signer struct {
	Secret     []byte
	Issuer     string
	Audience   string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func (s HS256Signer) NewAccessToken(userID string) (string, error) {
	now := time.Now()
	claims := Claims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    s.Issuer,
			Subject:   userID,
			Audience:  jwt.ClaimStrings{s.Audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(s.RefreshTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        generateJTI(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.Secret)
}

func (s HS256Signer) NewRefreshToken(userID string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    s.Issuer,
		Subject:   userID,
		Audience:  jwt.ClaimStrings{s.Audience},
		ExpiresAt: jwt.NewNumericDate(now.Add(s.RefreshTTL)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        generateJTI(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.Secret)
}

func generateJTI() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
