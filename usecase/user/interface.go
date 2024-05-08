package user

import _entities "booking-venue-api/entities/user"

type UserUseCaseInterface interface {
	CreateUser(request _entities.User) (_entities.User, error)
}
