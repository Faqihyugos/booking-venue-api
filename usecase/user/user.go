package user

import (
	"errors"

	"booking-venue-api/delivery/helper"
	_input "booking-venue-api/delivery/input"
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

func (uuc *UserUseCase) RegisterUser(input _input.RegisterUserInput) (_entities.User, error) {
	user := _entities.User{}

	user.Email = input.Email
	user.Username = input.Username
	user.Fullname = input.Fullname
	user.PhoneNumber = input.PhoneNumber

	password, err := helper.HashPassword(input.Password)
	if err != nil {
		return _entities.User{}, err
	}

	user.Password = password

	newUser, err := uuc.userRepository.Create(user)
	if err != nil {
		return _entities.User{}, err
	}
	return newUser, nil
}

func (uuc *UserUseCase) LoginUser(input _input.LoginInput) (string, error) {

	email := input.Email
	password := input.Password

	user, err := uuc.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", errors.New("no user found on that email")
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

func (uuc *UserUseCase) UpdateUser(ID int, inputData _input.UpdateUserInput) (_entities.User, error) {
	user, err := uuc.userRepository.FindByID(ID)
	if err != nil {
		return _entities.User{}, err
	}

	user.Email = inputData.Email
	user.Username = inputData.Username
	user.Fullname = inputData.Fullname
	user.PhoneNumber = inputData.PhoneNumber

	password, err := helper.HashPassword(inputData.Password)
	if err != nil {
		return _entities.User{}, err
	}

	user.Password = password

	user, err = uuc.userRepository.Update(user)
	if err != nil {
		return _entities.User{}, err
	}
	return user, nil
}

func (uuc *UserUseCase) GetUserByID(ID int) (_entities.User, error) {

	user, err := uuc.userRepository.FindByID(ID)
	if err != nil {
		return _entities.User{}, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (uuc *UserUseCase) DeleteUser(ID int) (_entities.User, error) {

	user, err := uuc.userRepository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found on with that ID")
	}

	deleteUser, err := uuc.userRepository.Delete(user)
	if err != nil {
		return deleteUser, err
	}

	return deleteUser, nil

}
