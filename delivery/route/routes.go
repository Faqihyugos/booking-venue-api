package route

import (
	_userHandler "booking-venue-api/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	e.POST("/api/register", uh.CreateUserHandler())
}
