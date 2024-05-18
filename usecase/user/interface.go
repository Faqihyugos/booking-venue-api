package user

import _entities "booking-venue-api/entities/user"

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
	LoginUser(email string, password string) (string, error)
	GetUserByID(id int) (_entities.User, error)
}
