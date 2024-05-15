package category

import _entities "booking-venue-api/entities"

type CategoryRepositoryInterface interface {
	FindAll() ([]_entities.Category, error)
	FindByID(id uint) (_entities.Category, error)
	Store(category _entities.Category) (_entities.Category, error)
	Update(category _entities.Category) (_entities.Category, error)
	Delete(category _entities.Category) (_entities.Category, error)
}
