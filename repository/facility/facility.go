package facility

import (
	_entities "booking-venue-api/entities"

	"gorm.io/gorm"
)

type FacilityRepository struct {
	DB *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) *FacilityRepository {
	return &FacilityRepository{DB: db}
}

func (fr *FacilityRepository) FindAll() ([]_entities.Facility, error) {
	var facilities []_entities.Facility
	err := fr.DB.Find(&facilities).Error
	if err != nil {
		return []_entities.Facility{}, err
	}
	return facilities, nil
}

func (fr *FacilityRepository) FindByID(id uint) (_entities.Facility, error) {
	var facility _entities.Facility
	err := fr.DB.First(&facility, id).Error
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}

func (fr *FacilityRepository) Store(facility _entities.Facility) (_entities.Facility, error) {
	err := fr.DB.Create(&facility).Error
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}

func (fr *FacilityRepository) Update(facility _entities.Facility) (_entities.Facility, error) {
	err := fr.DB.Save(&facility).Error
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}

func (fr *FacilityRepository) Delete(facility _entities.Facility) (_entities.Facility, error) {
	err := fr.DB.Delete(&facility).Error
	if err != nil {
		return _entities.Facility{}, err
	}
	return facility, nil
}
