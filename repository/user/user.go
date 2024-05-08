package user

import (
	_entities "booking-venue-api/entities/user"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(request _entities.User) (_entities.User, error) {
	err := ur.DB.Create(&request).Error
	if err != nil {
		return _entities.User{}, err
	}

	return request, nil
}

func (ur *UserRepository) GetByEmail(email string) (_entities.User, error) {
	var user _entities.User
	err := ur.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return _entities.User{}, err
	}

	return user, nil
}
