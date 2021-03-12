package repo

import (
	"context"

	"github.com/matheusmosca/nutcache/model"
	"gorm.io/gorm"
)

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

type EmployeeRepo interface {
	Create(context.Context, *model.Employee) error
	GetAll(context.Context) ([]*model.Employee, error)
	GetByID(context.Context, int) (*model.Employee, error)
	Update(context.Context, *model.Employee) error
	Delete(context.Context, int) error
}

func (r employeeRepo) Create(ctx context.Context, employee *model.Employee) error {
	err := r.db.Create(employee).Error
	if err != nil {
		return err
	}
	return nil
}

func (r employeeRepo) GetAll(ctx context.Context) ([]*model.Employee, error) {
	var employees []*model.Employee

	err := r.db.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (r employeeRepo) GetByID(ctx context.Context, ID int) (*model.Employee, error) {
	var employee model.Employee

	err := r.db.First(&employee, ID).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r employeeRepo) Update(ctx context.Context, employee *model.Employee) error {
	return r.db.Save(&employee).Error
}

func (r employeeRepo) Delete(ctx context.Context, ID int) error {
	return r.db.Delete(&model.Employee{}, ID).Error
}
