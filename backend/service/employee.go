package service

import "github.com/matheusmosca/nutcache/infrastructure/repo"

type employeeService struct {
	repo repo.EmployeeRepo
}

func NewEmployeeService(repo repo.EmployeeRepo) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

type EmployeeService interface {
	// CreateEmployee(context.Context, *model.Employee) error
}
