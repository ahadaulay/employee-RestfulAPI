package services

import (
	"catering-api/models/dto"
	"catering-api/repositories"
)

type EmployeeService interface {
	FindAll() ([]dto.EmployeeResponse, error)
	FindById(id uint) (dto.EmployeeResponse, error)
	Create(request dto.EmployeeRequest) (dto.EmployeeRequest, error)
	Update(request dto.EmployeeRequest, id uint) (dto.EmployeeRequest, error)
	Delete(id uint) error
}

type EmployeeImplementation struct {
	repository repositories.EmployeeRepository
}

func (service *EmployeeImplementation) FindAll() ([]dto.EmployeeResponse, error) {
	employee, err := service.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (service *EmployeeImplementation) FindById(id uint) (dto.EmployeeResponse, error) {
	employee, err := service.repository.FindById(id)

	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return employee, nil

}

func (service *EmployeeImplementation) Create(request dto.EmployeeRequest) (dto.EmployeeRequest, error) {
	employee, err := service.repository.Create(request)

	if err != nil {
		return dto.EmployeeRequest{}, err
	}

	return employee, nil

}

func (service *EmployeeImplementation) Update(request dto.EmployeeRequest, id uint) (dto.EmployeeRequest, error) {
	employee, err := service.repository.Update(request, id)

	if err != nil {
		return dto.EmployeeRequest{}, err
	}

	return employee, nil

}

func (service *EmployeeImplementation) Delete(id uint) error {
	err := service.repository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}

func NewEmployeeService(Employeerepo repositories.EmployeeRepository) EmployeeService {
	return &EmployeeImplementation{
		repository: Employeerepo,
	}
}
