package presenter

import (
	"time"
)

type EmployeeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Team      string    `json:"team"`
	CPF       string    `json:"cpf"`
	StartDate time.Time `json:"startDate"`
	BirthDate time.Time `json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type EmployeeRequest struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Team      string    `json:"team"`
	CPF       string    `json:"cpf"`
	StartDate time.Time `json:"startDate"`
	BirthDate time.Time `json:"birthDate"`
}

type EmployeeUpdate struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Team      string    `json:"team"`
	CPF       string    `json:"cpf"`
	StartDate time.Time `json:"startDate"`
	BirthDate time.Time `json:"birthDate"`
}
