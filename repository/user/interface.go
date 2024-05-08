package user

import _entities "booking-venue-api/entities/user"

type UserRepositoryInterface interface {
	Create(request _entities.User) (_entities.User, error)
}
