package category

import (
	"booking-venue-api/delivery/helper"
	_entities "booking-venue-api/entities"
	_categoryUseCase "booking-venue-api/usecase/category"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUseCase _categoryUseCase.CategoryUseCaseInterface
}

func NewCategoryHandler(u _categoryUseCase.CategoryUseCaseInterface) CategoryHandler {
	return CategoryHandler{
		categoryUseCase: u,
	}
}

// get all
func (ch *CategoryHandler) GetCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := ch.categoryUseCase.GetAllCategory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to get categories"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully get categories", categories))
	}
}

// create
func (ch *CategoryHandler) CreateCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var category _entities.Category
		err := c.Bind(&category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to bind data"))
		}
		response, err := ch.categoryUseCase.CreateCategory(category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to create category"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully create category", response))
	}
}

// update
func (ch *CategoryHandler) UpdateCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var category _entities.Category
		err := c.Bind(&category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to bind data"))
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to convert id"))
		}
		response, err := ch.categoryUseCase.UpdateCategory(uint(id), category)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to update category"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully update category", response))
	}
}

// delete
func (ch *CategoryHandler) DeleteCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to convert id"))
		}
		err = ch.categoryUseCase.DeleteCategory(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to delete category"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully delete category"))
	}
}
