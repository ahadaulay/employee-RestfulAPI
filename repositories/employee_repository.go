package repositories

import (
	"catering-api/models/dto"
	"catering-api/models/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindAll() ([]dto.EmployeeResponse, error)
	FindById(id uint) (dto.EmployeeResponse, error)
	Create(request dto.EmployeeRequest) (dto.EmployeeRequest, error)
	Update(request dto.EmployeeRequest, id uint) (dto.EmployeeRequest, error)
	Delete(id uint) error
}

type EmployeeImplementation struct {
	db *gorm.DB
}

func (repository *EmployeeImplementation) FindAll() ([]dto.EmployeeResponse, error) {
	employeeModel := []model.Employee{}

	result := repository.db.Where("deleted_at IS NULL").Find(&employeeModel)

	if result.Error != nil {
		return nil, result.Error
	}

	var employee []dto.EmployeeResponse

	if err := copier.Copy(&employee, &employeeModel); err != nil {
		return nil, err
	}

	return employee, nil

}

func (repository *EmployeeImplementation) FindById(id uint) (dto.EmployeeResponse, error) {
	employeeModel := model.Employee{}

	result := repository.db.Where("deleted_at IS NULL").First(&employeeModel, id)

	if result.Error != nil {
		return dto.EmployeeResponse{}, result.Error
	}

	var employee dto.EmployeeResponse

	if err := copier.Copy(&employee, &employeeModel); err != nil {
		return dto.EmployeeResponse{}, err
	}

	return employee, nil

}

func (repository *EmployeeImplementation) Create(request dto.EmployeeRequest) (dto.EmployeeRequest, error) {
	var employeeModel model.Employee

	if err := copier.Copy(&employeeModel, &request); err != nil {
		return dto.EmployeeRequest{}, err
	}

	result := repository.db.Create(&employeeModel)

	if result.Error != nil {
		return dto.EmployeeRequest{}, result.Error
	}

	return request, nil
}

func (repository *EmployeeImplementation) Update(request dto.EmployeeRequest, id uint) (dto.EmployeeRequest, error) {

	result := repository.db.Table("employees").Where("id = ?", id).Updates(request)

	if result.Error != nil {
		return dto.EmployeeRequest{}, result.Error
	}

	return request, nil
}

func (repository *EmployeeImplementation) Delete(id uint) error {
	result := repository.db.Delete(&model.Employee{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &EmployeeImplementation{db: db}
}
