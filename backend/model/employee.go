package model

import (
	"time"
)

type Employee struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null; unique" json:"email"`
	Gender    string    `gorm:"not null" json:"gender"`
	Team      string    `gorm:"null" json:"team"`
	CPF       string    `gorm:"unique;not null" json:"cpf"`
	StartDate time.Time `gorm:"not null" json:"startDate"`
	BirthDate time.Time `gorm:"not null" json:"birthDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewEmployee(name, email, gender, team, CPF string, birthDate, startDate time.Time) (*Employee, error) {
	e := Employee{
		Name:      name,
		Email:     email,
		Gender:    gender,
		Team:      team,
		CPF:       CPF,
		BirthDate: birthDate,
		StartDate: startDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := e.Validate()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Employee) Validate() error {
	switch {
	case e.Name == "":
		return ErrInvalidName
	case e.Email == "":
		return ErrInvalidEmail
	case len(e.CPF) != 11:
		return ErrInvalidCPF
	case e.Gender == "" || (e.Gender != "Female" && e.Gender != "Male" && e.Gender != "Other"):
		return ErrInvalidGender
	case e.Team != "" && (e.Team != "Backend" && e.Team != "Frontend" && e.Team != "Mobile"):
		return ErrInvalidTeam
	default:
		return nil
	}
}
