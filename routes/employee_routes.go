package routes

import (
	"catering-api/controllers"
	"catering-api/repositories"
	"catering-api/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EmployeeRoutes(e *echo.Echo, db *gorm.DB) {
	employeeRepository := repositories.NewEmployeeRepository(db)

	employeeService := services.NewEmployeeService(employeeRepository)

	employeeController := controllers.EmployeeController{
		EmployeeService: employeeService,
	}

	employeeGroup := e.Group("api/v1/employee")

	employeeGroup.GET("", employeeController.FindAll)
	employeeGroup.GET("/:id", employeeController.FindById)
	employeeGroup.POST("", employeeController.Create)
	employeeGroup.PUT("/:id", employeeController.Update)
	employeeGroup.DELETE("/:id", employeeController.Delete)

}
