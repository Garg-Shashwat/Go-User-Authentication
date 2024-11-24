package transport

import (
	"net/http"

	loginControllers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/login/controllers"
	logoutControllers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/logout/controllers"
	refreshControllers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/refresh_token/controllers"
	registerControllers "github.com/Garg-Shashwat/Go-User-Authentication/root/routes/register/controllers"
	"github.com/Garg-Shashwat/Go-User-Authentication/root/transport/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// addRoutes adds all routes to server
func addRoutes(router *echo.Echo) {
	router.Pre(echoMiddleware.RemoveTrailingSlash())

	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Pong!")
	})

	router.POST("/register", registerControllers.RegisterUser)
	router.POST("/login", loginControllers.LoginUser)

	router.Use(middleware.AuthorizeUser)
	router.GET("/refresh_token", refreshControllers.RefreshToken)
	router.GET("/logout", logoutControllers.LogoutUser)
}
