package transport

import (
	"net/http"

	"github.com/Garg-Shashwat/Go-User-Authentication/root/routes/register/controllers"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// addRoutes adds all routes to server
func addRoutes(router *echo.Echo) {
	router.Pre(echoMiddleware.RemoveTrailingSlash())

	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Pong!")
	})

	router.POST("/register", controllers.RegisterUser)
}
