package category

import (
	_entities "booking-venue-api/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) FindAll() ([]_entities.Category, error) {
	var categories []_entities.Category
	err := cr.DB.Find(&categories).Error
	if err != nil {
		return []_entities.Category{}, err
	}
	return categories, nil
}

func (cr *CategoryRepository) FindByID(id uint) (_entities.Category, error) {
	var category _entities.Category
	err := cr.DB.Where("id = ?", id).Take(&category).Error
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepository) Store(category _entities.Category) (_entities.Category, error) {
	err := cr.DB.Create(&category).Error
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepository) Update(category _entities.Category) (_entities.Category, error) {
	err := cr.DB.Save(&category).Error
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepository) Delete(category _entities.Category) (_entities.Category, error) {
	err := cr.DB.Delete(&category).Error
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}
