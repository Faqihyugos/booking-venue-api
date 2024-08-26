package facility

import _entities "booking-venue-api/entities"

type IFacilityRepository interface {
	FindAll() ([]_entities.Facility, error)
	FindByID(id uint) (_entities.Facility, error)
	Store(facility _entities.Facility) (_entities.Facility, error)
	Update(facility _entities.Facility) (_entities.Facility, error)
	Delete(facility _entities.Facility) (_entities.Facility, error)
}
