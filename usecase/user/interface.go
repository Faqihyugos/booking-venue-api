package user

import (
	_input "booking-venue-api/delivery/input"
	_entities "booking-venue-api/entities"
)

type UserUseCaseInterface interface {
	RegisterUser(input _input.RegisterUserInput) (_entities.User, error)
	LoginUser(input _input.LoginInput) (string, error)
	UpdateUser(ID int, inputData _input.UpdateUserInput) (_entities.User, error)
	GetUserByID(ID int) (_entities.User, error)
	DeleteUser(ID int) (_entities.User, error)
}
