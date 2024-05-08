package user

import (
	"errors"

	"booking-venue-api/delivery/helper"
	_entities "booking-venue-api/entities/user"
	_userRepository "booking-venue-api/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) CreateUser(request _entities.User) (_entities.User, error) {
	password, err := helper.HashPassword(request.Password)
	request.Password = password
	users, err := uuc.userRepository.Create(request)

	if request.Fullname == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Email == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Password == "" {
		return users, errors.New("Can't be empty")
	}
	if request.PhoneNumber == "" {
		return users, errors.New("Can't be empty")
	}
	if request.Username == "" {
		return users, errors.New("Can't be empty")
	}

	return users, err
}