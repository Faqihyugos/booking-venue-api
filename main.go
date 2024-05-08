package main

import (
	"fmt"
	"log"
	"net/http"

	"booking-venue-api/config"

	_userHandler "booking-venue-api/delivery/handler/user"
	_middleware "booking-venue-api/delivery/middleware"
	_route "booking-venue-api/delivery/route"
	_userRepository "booking-venue-api/repository/user"
	_userUseCase "booking-venue-api/usecase/user"
	_utils "booking-venue-api/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.NewConfig()
	db := _utils.InitDB(config.DB)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(_middleware.CustomLogger())

	_route.RegisterUserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.HTTP.Port)))
}
