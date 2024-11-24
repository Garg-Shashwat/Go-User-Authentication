package middleware

import (
	"net/http"
	"strings"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/utils"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthorizeUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Token")
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token Format")
		}

		tokenString := bearerToken[1]

		token, err := utils.VerifyJWTToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(401, "Invalid Token Claims")
		}

		ctx.Set("userID", claims["sub"].(uint))

		return next(ctx)
	}
}
