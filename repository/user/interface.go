package user

import _entities "booking-venue-api/entities"

type UserRepositoryInterface interface {
	Create(user _entities.User) (_entities.User, error)
	FindByEmail(email string) (_entities.User, error)
	FindByID(ID int) (_entities.User, error)
	Update(user _entities.User) (_entities.User, error)
	Delete(user _entities.User) (_entities.User, error)
}
