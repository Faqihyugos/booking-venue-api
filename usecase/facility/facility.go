package facility

import (
	_entities "booking-venue-api/entities"
	_facilityRepository "booking-venue-api/repository/facility"
)

type FacilityUseCase struct {
	FacilityRepository _facilityRepository.IFacilityRepository
}

func NewFacilityUseCase(facilityRepo _facilityRepository.IFacilityRepository) *FacilityUseCase {
	return &FacilityUseCase{FacilityRepository: facilityRepo}
}

func (fuc *FacilityUseCase) GetFacilities() ([]_entities.Facility, error) {
	facilities, err := fuc.FacilityRepository.FindAll()
	if err != nil {
		return []_entities.Facility{}, err
	}
	return facilities, nil
}

func (fuc *FacilityUseCase) CreateFacility(request _entities.Facility) (_entities.Facility, error) {
	facility, err := fuc.FacilityRepository.Store(request)
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}

func (fuc *FacilityUseCase) UpdateFacility(id uint, request _entities.Facility) (_entities.Facility, error) {
	// find by id form repo
	facility, err := fuc.FacilityRepository.FindByID(id)
	if err != nil {
		return _entities.Facility{}, err
	}

	// update
	facility.Name = request.Name
	facility.IconName = request.IconName
	facility, err = fuc.FacilityRepository.Update(facility)
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}

func (fuc *FacilityUseCase) DeleteFacility(id uint) error {
	facility, err := fuc.FacilityRepository.FindByID(id)
	if err != nil {
		return err
	}

	_, err = fuc.FacilityRepository.Delete(facility)
	if err != nil {
		return err
	}
	return nil
}
