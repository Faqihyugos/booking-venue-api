package user

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"booking-venue-api/delivery/helper"
	_input "booking-venue-api/delivery/input"
	_middlewares "booking-venue-api/delivery/middleware"
	_userUseCase "booking-venue-api/usecase/user"
	"booking-venue-api/utils"

	"github.com/go-playground/validator/v10"
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

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (uh *UserHandler) RegisterUserHandler() echo.HandlerFunc {

	return func(c echo.Context) error {
		var input _input.RegisterUserInput
		errBind := c.Bind(&input)
		if errBind != nil {
			errors := utils.FormatValidationError(errBind)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed(errorMessage))
		}

		// Validate the input
		errValidate := validate.Struct(input)
		if errValidate != nil {
			errors := utils.FormatValidationError(errValidate)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errorMessage))
		}

		_, err := uh.userUseCase.RegisterUser(input)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully registered"))
	}
}

func (uh *UserHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var login _input.LoginInput
		err := c.Bind(&login)
		if err != nil {
			errors := utils.FormatValidationError(err)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed(errorMessage))
		}

		// Validate the input
		errValidate := validate.Struct(login)
		if errValidate != nil {
			errors := utils.FormatValidationError(errValidate)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errorMessage))
		}

		token, errorLogin := uh.userUseCase.LoginUser(login)
		if errorLogin != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(fmt.Sprintf("%v", errorLogin)))
		}
		responseToken := map[string]interface{}{
			"token": token,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully logged in", responseToken))
	}
}

func (uh *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var updateRequest _input.UpdateUserInput

		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			fmt.Println("idToken : ", idToken)
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized"))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized"))
		}
		err := c.Bind(&updateRequest)
		if err != nil {
			errors := utils.FormatValidationError(err)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusUnprocessableEntity, helper.ResponseFailed(errorMessage))
		}

		// Validate the input
		errValidate := validate.Struct(updateRequest)
		if errValidate != nil {
			errors := utils.FormatValidationError(errValidate)
			errorMessage := strings.Join(errors, ", ")
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed(errorMessage))
		}

		users, err := uh.userUseCase.UpdateUser(userId, updateRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully update user", users))
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		// check login status
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			fmt.Println("idToken : ", idToken)
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized"))
		}

		userId, _ := strconv.Atoi(c.Param("userId"))
		// check authorization
		if idToken != userId {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Unauthorized"))
		}
		_, err := uh.userUseCase.DeleteUser(userId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed(err.Error()))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully delete user"))
	}
}
