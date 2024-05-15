package category

import (
	_entities "booking-venue-api/entities"
	_categoryRepository "booking-venue-api/repository/category"
)

type CategoryUseCase struct {
	categoryRepository _categoryRepository.CategoryRepositoryInterface
}

func NewCategoryUseCase(categoryRepo _categoryRepository.CategoryRepositoryInterface) *CategoryUseCase {
	return &CategoryUseCase{categoryRepository: categoryRepo}
}
func (cuu *CategoryUseCase) GetAllCategory() ([]_entities.Category, error) {
	categories, err := cuu.categoryRepository.FindAll()
	if err != nil {
		return []_entities.Category{}, err
	}
	return categories, nil
}

func (cuu *CategoryUseCase) CreateCategory(request _entities.Category) (_entities.Category, error) {
	category, err := cuu.categoryRepository.Store(request)
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}
func (cuu *CategoryUseCase) UpdateCategory(id uint, request _entities.Category) (_entities.Category, error) {
	// find by id form repo
	category, err := cuu.categoryRepository.FindByID(id)
	if err != nil {
		return _entities.Category{}, err
	}
	// update category from result find by id and update by request
	category.Name = request.Name
	category.IconName = request.IconName
	category, err = cuu.categoryRepository.Update(category)
	if err != nil {
		return _entities.Category{}, err
	}
	return category, nil
}

func (cuu *CategoryUseCase) DeleteCategory(id uint) error {
	category, err := cuu.categoryRepository.FindByID(id)
	if err != nil {
		return err
	}
	_, err = cuu.categoryRepository.Delete(category)
	if err != nil {
		return err
	}
	return nil
}
