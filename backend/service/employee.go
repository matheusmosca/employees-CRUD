package service

import (
	"context"
	"time"

	"github.com/matheusmosca/nutcache/infrastructure/repo"
	"github.com/matheusmosca/nutcache/model"
)

type employeeService struct {
	repo repo.EmployeeRepo
}

func NewEmployeeService(repo repo.EmployeeRepo) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

type EmployeeService interface {
	Create(context.Context, string, string, string, string, string, time.Time, time.Time) error
	GetAll(context.Context) ([]*model.Employee, error)
	Update(context.Context, int, string, string, string, string, string, time.Time, time.Time) error
	GetByID(context.Context, int) (*model.Employee, error)
	Delete(context.Context, int) error
}

func (e *employeeService) Create(ctx context.Context, name, gender, email, CPF, team string, startDate, birthDate time.Time) error {
	employee, err := model.NewEmployee(name, email, gender, team, CPF, birthDate, startDate)

	if err != nil {
		return err
	}
	return e.repo.Create(ctx, employee)
}

func (e *employeeService) Update(ctx context.Context, ID int, name, gender, email, CPF, team string, startDate, birthDate time.Time) error {
	_, err := model.NewEmployee(name, email, gender, team, CPF, birthDate, startDate)

	if err != nil {
		return err
	}

	employee, err := e.repo.GetByID(ctx, ID)
	if err != nil {
		return err
	}

	employee.Name = name
	employee.Gender = gender
	employee.Email = email
	employee.CPF = CPF
	employee.Team = team
	employee.StartDate = startDate
	employee.BirthDate = birthDate

	return e.repo.Update(ctx, employee)
}

func (e *employeeService) GetAll(ctx context.Context) ([]*model.Employee, error) {
	return e.repo.GetAll(ctx)
}

func (e *employeeService) GetByID(ctx context.Context, ID int) (*model.Employee, error) {
	return e.repo.GetByID(ctx, ID)
}

func (e *employeeService) Delete(ctx context.Context, ID int) error {
	return e.repo.Delete(ctx, ID)
}
