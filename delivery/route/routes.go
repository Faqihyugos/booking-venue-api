package route

import (
	_categoryHandler "booking-venue-api/delivery/handler/category"
	_userHandler "booking-venue-api/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func AuthPath(e *echo.Echo, uh _userHandler.UserHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.POST("/register", uh.CreateUserHandler())
	apiGroup.POST("/login", uh.LoginHandler())
}

func CategoryPath(e *echo.Echo, uh _categoryHandler.CategoryHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/categories", uh.GetCategoryHandler())
	apiGroup.POST("/categories", uh.CreateCategoryHandler())
	apiGroup.PUT("/categories/:id", uh.UpdateCategoryHandler())
	apiGroup.DELETE("/categories/:id", uh.DeleteCategoryHandler())
}
