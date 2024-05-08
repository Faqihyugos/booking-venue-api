package user

import (
	"errors"

	"booking-venue-api/delivery/helper"
	"booking-venue-api/delivery/middleware"
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

func (uuc *UserUseCase) LoginUser(email string, password string) (string, error) {
	user, err := uuc.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !helper.CheckPassHash(password, user.Password) {
		return "", errors.New("Wrong password")
	}
	token, err := middleware.CreateToken(int(user.ID), user.Username)
	if err != nil {
		return "", err
	}

	if email == "" {
		return "", errors.New("Email can't be empty")
	}
	if password == "" {
		return "", errors.New("Password can't be empty")
	}

	return token, nil
}
