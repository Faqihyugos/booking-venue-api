package facility

import (
	_entities "booking-venue-api/entities"
)

type IFacilityUsecase interface {
	GetFacilities() ([]_entities.Facility, error)
	CreateFacility(request _entities.Facility) (_entities.Facility, error)
	UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, error)
	DeleteFacility(id uint) error
}
