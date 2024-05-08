package route

import (
	_userHandler "booking-venue-api/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func AuthPath(e *echo.Echo, uh _userHandler.UserHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.POST("/register", uh.CreateUserHandler())
	apiGroup.POST("/login", uh.LoginHandler())
}
