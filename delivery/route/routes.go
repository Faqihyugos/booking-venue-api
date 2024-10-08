package route

import (
	_categoryHandler "booking-venue-api/delivery/handler/category"
	_facilityHandler "booking-venue-api/delivery/handler/facility"
	_userHandler "booking-venue-api/delivery/handler/user"
	_middlewares "booking-venue-api/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func AuthPath(e *echo.Echo, uh _userHandler.UserHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.POST("/register", uh.RegisterUserHandler())
	apiGroup.POST("/login", uh.LoginHandler())
}

func UserPath(e *echo.Echo, uh _userHandler.UserHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.PUT("/users/:userId", uh.UpdateUser(), _middlewares.JWTMiddleware())
	apiGroup.DELETE("/users/:userId", uh.DeleteUser(), _middlewares.JWTMiddleware())
	apiGroup.GET("/users/:userId", uh.GetUserByID(), _middlewares.JWTMiddleware())
}

func CategoryPath(e *echo.Echo, uh _categoryHandler.CategoryHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/categories", uh.GetCategoryHandler())
	apiGroup.POST("/categories", uh.CreateCategoryHandler())
	apiGroup.PUT("/categories/:id", uh.UpdateCategoryHandler())
	apiGroup.DELETE("/categories/:id", uh.DeleteCategoryHandler())
}

func FacilityPath(e *echo.Echo, uh _facilityHandler.FacilityHandler) {
	apiGroup := e.Group("/api/v1")
	apiGroup.GET("/facilities", uh.GetFacilityHandler())
	apiGroup.POST("/facilities", uh.CreateFacilityHandler())
	apiGroup.PUT("/facilities/:id", uh.UpdateFacilityHandler())
	apiGroup.DELETE("/facilities/:id", uh.DeleteFacilityHandler())
}
