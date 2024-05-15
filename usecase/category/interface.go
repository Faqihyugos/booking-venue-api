package category

import _entities "booking-venue-api/entities"

type CategoryUseCaseInterface interface {
	GetAllCategory() ([]_entities.Category, error)
	CreateCategory(request _entities.Category) (_entities.Category, error)
	UpdateCategory(id uint, request _entities.Category) (_entities.Category, error)
	DeleteCategory(id uint) error
}
