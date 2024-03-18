package tokenizer

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken            = errors.New("token is invalid")
	ErrTokenExpired            = errors.New("token is expired")
	JWTErrorTokenExpiredString = "token has invalid claims: token is expire"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func VerifyToken(token string, secret []byte) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return secret, nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		switch err.Error() {
		case JWTErrorTokenExpiredString:
			fallthrough
		case jwt.ErrTokenExpired.Error():
			return nil, ErrTokenExpired
		case jwt.ErrTokenInvalidClaims.Error():
			return nil, ErrInvalidToken
		default:
			return nil, ErrInvalidToken
		}
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func NewPayload(user_uuid uuid.UUID, name string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:               tokenID,
		UUID:             user_uuid,
		Name:             name,
		IssuedAt:         time.Now(),
		ExpiredAt:        time.Now().Add(duration),
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration))},
	}

	return payload, nil
}

func CreateToken(payload *Payload, secret []byte) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}
