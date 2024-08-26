package user

import (
	_entities "booking-venue-api/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.LoginUser("faq@email.com", "wrong_password")
		assert.EqualError(t, err, "Wrong password")
		assert.Equal(t, "", data)
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

// GetByID implements user.UserRepositoryInterface.
func (m mockUserRepository) GetByID(id int) (_entities.User, error) {
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
		Email:       "faq@email.com",
		Password:    "12345",
		Username:    "faq",
		Fullname:    "faq",
		PhoneNumber: "082",
	}, nil
}

// GetByEmail implements user.UserRepositoryInterface.
func (m mockUserRepository) GetByEmail(email string) (_entities.User, error) {
	return _entities.User{
		Email:       "faq@email.com",
		Password:    "12345",
		Username:    "faq",
		Fullname:    "faq",
		PhoneNumber: "082",
	}, nil
}
