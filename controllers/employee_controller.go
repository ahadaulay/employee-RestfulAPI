package controllers

import (
	"catering-api/helpers"
	"catering-api/models/dto"
	"catering-api/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	EmployeeService services.EmployeeService
}

func (controller *EmployeeController) FindAll(c echo.Context) error {
	employee, err := controller.EmployeeService.FindAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("fail get all employee", err))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success get all employee", employee))

}

func (controller *EmployeeController) FindById(c echo.Context) error {
	employeeID := c.Param("id")
	employeeIDInt, err := strconv.Atoi(employeeID)
	employeeIDUint := uint(employeeIDInt)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id", err))
	}

	employee, err := controller.EmployeeService.FindById(employeeIDUint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("fail get employee", err))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success get  employee", employee))

}

func (controller *EmployeeController) Create(c echo.Context) error {
	var employee dto.EmployeeRequest

	err := c.Bind(&employee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed to bind data", err))
	}

	employee, err = controller.EmployeeService.Create(employee)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("fail create employee", err))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessResponse("success create employee", employee))
}

func (controller *EmployeeController) Update(c echo.Context) error {
	var employee dto.EmployeeRequest
	employeeID := c.Param("id")
	employeeIDInt, err := strconv.Atoi(employeeID)
	employeeIDUint := uint(employeeIDInt)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id", err))
	}

	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid request data", err))
	}

	if _, err := controller.EmployeeService.Update(employee, employeeIDUint); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Failed to update employee", err))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("employee updated successfully", employee))

}

func (controller *EmployeeController) Delete(c echo.Context) error {
	employeeID := c.Param("id")
	employeeIDInt, err := strconv.Atoi(employeeID)
	employeeIDUint := uint(employeeIDInt)

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id", err))
	}

	err = controller.EmployeeService.Delete(employeeIDUint)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("fail delete employee", err))
	}

	return c.JSON(http.StatusOK, helpers.SuccessResponse("success delete employee", nil))

}
