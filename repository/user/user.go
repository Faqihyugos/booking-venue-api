package user

import (
	_entities "booking-venue-api/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user _entities.User) (_entities.User, error) {
	err := ur.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) FindByEmail(email string) (_entities.User, error) {
	var user _entities.User
	err := ur.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) FindByID(ID int) (_entities.User, error) {
	var user _entities.User

	err := ur.DB.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Update(user _entities.User) (_entities.User, error) {
	err := ur.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Delete(user _entities.User) (_entities.User, error) {
	err := ur.DB.Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
