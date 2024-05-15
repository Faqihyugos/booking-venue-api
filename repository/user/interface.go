package user

import _entities "booking-venue-api/entities"

type UserRepositoryInterface interface {
	Create(request _entities.User) (_entities.User, error)
	GetByEmail(email string) (_entities.User, error)
}
