package facility

import (
	"booking-venue-api/delivery/helper"
	_entities "booking-venue-api/entities"
	_facilityUseCase "booking-venue-api/usecase/facility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FacilityHandler struct {
	facilityUseCase _facilityUseCase.IFacilityUsecase
}

func NewFacilityHandler(facilityUseCase _facilityUseCase.IFacilityUsecase) FacilityHandler {
	return FacilityHandler{facilityUseCase}
}

// get all Facility
func (fh *FacilityHandler) GetFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		facilities, err := fh.facilityUseCase.GetFacilities()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to get facilities"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully get facilities", facilities))
	}
}

// create
func (fh *FacilityHandler) CreateFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var facility _entities.Facility
		err := c.Bind(&facility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to bind data"))
		}
		response, err := fh.facilityUseCase.CreateFacility(facility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to create facility"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully create facility", response))
	}
}

// update
func (fh *FacilityHandler) UpdateFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var facility _entities.Facility
		err := c.Bind(&facility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to bind data"))
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to convert id"))
		}
		response, err := fh.facilityUseCase.UpdateFacility(uint(id), facility)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to update facility"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("Successfully update facility", response))
	}
}

// delete
func (fh *FacilityHandler) DeleteFacilityHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to convert id"))
		}
		err = fh.facilityUseCase.DeleteFacility(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Error to delete facility"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("Successfully delete facility"))
	}
}
