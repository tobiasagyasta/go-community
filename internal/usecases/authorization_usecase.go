package usecases

// NOT USED
import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Authorization interface {
	Generate(accountNumber string, userId int) (string, error)
	Validate(token string) (*jwt.Token, error)
}

type Auth struct {
	secret string
	td     int
}

func NewAuthorization(secret string, td int) (*Auth, error) {
	if secret == "" {
		return nil, errors.New("empty signing key")
	}

	if td == 0 {
		return nil, errors.New("empty duration")
	}

	return &Auth{secret: secret, td: td}, nil
}

func (a *Auth) Generate(accountNumber string, userId int) (string, error) {
	expiry := time.Now().Add(time.Hour * time.Duration(a.td)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":            userId,
		"accountNumber": accountNumber,
		"createdAt":     time.Now().Unix(),
		"expiredAt":     expiry,
	})

	return token.SignedString([]byte(a.secret))
}
