package main

import (
	"fmt"
	"log"
	"net/http"

	"booking-venue-api/config"

	_categoryHandler "booking-venue-api/delivery/handler/category"
	_facilityHandler "booking-venue-api/delivery/handler/facility"
	_userHandler "booking-venue-api/delivery/handler/user"
	_middleware "booking-venue-api/delivery/middleware"
	_route "booking-venue-api/delivery/route"
	_categoryRepository "booking-venue-api/repository/category"
	_facilityRepository "booking-venue-api/repository/facility"
	_userRepository "booking-venue-api/repository/user"
	_categoryUseCase "booking-venue-api/usecase/category"
	_facilityUseCase "booking-venue-api/usecase/facility"
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

	categoryRepo := _categoryRepository.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryHandler := _categoryHandler.NewCategoryHandler(categoryUseCase)

	facilityRepo := _facilityRepository.NewFacilityRepository(db)
	facilityUseCase := _facilityUseCase.NewFacilityUseCase(facilityRepo)
	facilityHandler := _facilityHandler.NewFacilityHandler(facilityUseCase)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(_middleware.CustomLogger())

	_route.AuthPath(e, userHandler)
	_route.UserPath(e, userHandler)
	_route.CategoryPath(e, categoryHandler)
	_route.FacilityPath(e, facilityHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.HTTP.Port)))
}
