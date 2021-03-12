package repo

import (
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
	// CreateEmployee(context.Context, *model.Employee) error
}

// func (r repo) CreateEmployee(ctx context.Context, employee *model.Employee) error {
// 	r.db.Create()
// }
