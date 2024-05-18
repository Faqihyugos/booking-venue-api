package user

import (
	"fmt"
	"net/http"
	"strconv"

	"booking-venue-api/delivery/helper"
	_entities "booking-venue-api/entities/user"
	_userUseCase "booking-venue-api/usecase/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(u _userUseCase.UserUseCaseInterface) UserHandler {
	return UserHandler{
		userUseCase: u,
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {

		var param _entities.User

		errBind := c.Bind(&param)
		if errBind != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error binding data"))
		}
		_, err := uh.userUseCase.CreateUser(param)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Register failed"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully registered"))
	}
}

func (uh *UserHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login _entities.User
		err := c.Bind(&login)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to bind data"))
		}
		token, errorLogin := uh.userUseCase.LoginUser(login.Email, login.Password)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully logged in", responseToken))
	}
}

func (uh *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user _entities.User
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Error to bind data"))
		}
		user, err = uh.userUseCase.GetUserByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", err)))
		}
		responseUser := map[string]interface{}{
			"user": map[string]interface{}{
				"fullname": user.Fullname,
				"username": user.Username,
				"email": user.Email,
				"phone": user.PhoneNumber,
			},
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Get user successfully", responseUser))
	}
}