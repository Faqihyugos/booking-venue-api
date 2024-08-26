package user

import (
	_input "booking-venue-api/delivery/input"
	_entities "booking-venue-api/entities"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		loginInput := _input.LoginInput{
			Email:    "faq@email.com",
			Password: "12345",
		}
		data, err := userUseCase.LoginUser(loginInput)
		assert.Nil(t, err)
		assert.Equal(t, _entities.User{
			Email:       "faq@email.com",
			Password:    "12345",
			Username:    "faq",
			Fullname:    "faq",
			PhoneNumber: "082",
		}, data)
	})

	t.Run("TestLoginWrongPassword", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		loginInput := _input.LoginInput{
			Email:    "faq@email.com",
			Password: "wrong_password",
		}
		data, err := userUseCase.LoginUser(loginInput)
		assert.EqualError(t, err, "Wrong password")
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestRegister(t *testing.T) {
	t.Run("TestRegisterSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{
			Email:       "faq@email.com",
			Password:    "12345",
			Username:    "faq",
			Fullname:    "faq",
			PhoneNumber: "082",
		})
		assert.Nil(t, err)
		assert.Equal(t, _entities.User{
			Email:       "faq@email.com",
			Password:    "12345",
			Username:    "faq",
			Fullname:    "faq",
			PhoneNumber: "082",
		}, data)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetUserByID(1)
		assert.Nil(t, err)
		assert.Equal(t, _entities.User{
			Email:       "faq@email.com",
			Username:    "faq",
			Fullname:    "faq",
			PhoneNumber: "082",
		}, data)
	})
}

type mockUserRepository struct{}

// FindByID implements user.UserRepositoryInterface.
func (m mockUserRepository) FindByID(id int) (_entities.User, error) {
	return _entities.User{
		Email:       "faq@email.com",
		Username:    "faq",
		Fullname:    "faq",
		PhoneNumber: "082",
	}, nil
}

// Create implements user.UserRepositoryInterface.
func (m mockUserRepository) Create(request _entities.User) (_entities.User, error) {
	return _entities.User{
		Email:       request.Email,
		Password:    request.Password,
		Username:    request.Username,
		Fullname:    request.Fullname,
		PhoneNumber: request.PhoneNumber,
	}, nil
}

// FindByEmail implements user.UserRepositoryInterface.
func (m mockUserRepository) FindByEmail(email string) (_entities.User, error) {
	if email == "faq@email.com" {
		return _entities.User{
			Email:       "faq@email.com",
			Password:    "12345",
			Username:    "faq",
			Fullname:    "faq",
			PhoneNumber: "082",
		}, nil
	}
	return _entities.User{}, errors.New("User not found")
}

// GetByEmail implements user.UserRepositoryInterface.
func (m mockUserRepository) GetByEmail(email string) (_entities.User, error) {
	return m.FindByEmail(email)
}

// Update implements user.UserRepositoryInterface.
func (m mockUserRepository) Update(user _entities.User) (_entities.User, error) {
	return user, nil
}

// Delete implements user.UserRepositoryInterface.
func (m mockUserRepository) Delete(user _entities.User) (_entities.User, error) {
	return user, nil
}
