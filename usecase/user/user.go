package user

import (
	"errors"

	"booking-venue-api/delivery/helper"
	"booking-venue-api/delivery/middleware"
	_entities "booking-venue-api/entities"
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

	if request.Fullname == "" || request.Fullname == " " {
		return _entities.User{}, errors.New("fullname can't be empty")
	}
	if request.Email == "" || request.Email == " " {
		return _entities.User{}, errors.New("email can't be empty")
	}
	if request.Password == "" || request.Password == " " {
		return _entities.User{}, errors.New("password can't be empty")
	}
	if request.PhoneNumber == "" || request.PhoneNumber == " " {
		return _entities.User{}, errors.New("phone number can't be empty")
	}
	if request.Username == "" || request.Username == " " {
		return _entities.User{}, errors.New("username can't be empty")
	}

	password, err := helper.HashPassword(request.Password)
	if err != nil {
		return _entities.User{}, err
	}
	request.Password = password
	users, err := uuc.userRepository.Create(request)

	return users, err
}

func (uuc *UserUseCase) LoginUser(email string, password string) (string, error) {
	if email == "" || email == " " {
		return "", errors.New("email can't be empty")
	}
	if password == "" || password == " " {
		return "", errors.New("password can't be empty")
	}

	user, err := uuc.userRepository.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !helper.CheckPassHash(password, user.Password) {
		return "", errors.New("wrong password")
	}
	token, err := middleware.CreateToken(int(user.ID), user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
