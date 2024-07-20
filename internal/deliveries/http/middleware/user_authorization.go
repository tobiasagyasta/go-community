package middleware

import (
	"go-community/internal/config"
	"go-community/internal/deliveries/http/common/response"
	"go-community/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type jwtClaims struct {
	AccountNumber string `json:"accountNumber"`
	jwt.RegisteredClaims
}

func JWTMiddleware(config *config.Configuration) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			header := ctx.Request().Header.Get("Authorization")
			if header == "" {
				return response.Error(ctx, models.ErrorEmptyToken)
			}

			tokenString := header[len("Bearer "):]
			token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (sec interface{}, err error) {
				if config.Auth.BearerSecret == "" {
					return nil, err
				}
				return []byte(config.Auth.BearerSecret), nil
			})

			if err != nil {
				if err.Error() == "token has invalid claims: token is expired" {
					return response.Error(ctx, models.ErrorExpiredToken)
				}
				return response.Error(ctx, models.ErrorInvalidToken)
			}

			claims, ok := token.Claims.(*jwtClaims)
			if !ok || !token.Valid {
				return response.Error(ctx, models.ErrorInvalidToken)
			}

			if claims.ExpiresAt.Time.Before(time.Now()) {
				return response.Error(ctx, models.ErrorExpiredToken)
			}

			ctx.Set("accountNumber", claims.AccountNumber)
			return next(ctx)
		}
	}
}
