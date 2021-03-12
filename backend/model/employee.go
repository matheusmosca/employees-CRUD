package model

import (
	"time"
)

type Employee struct {
	ID        int       `gorm:primaryKey`
	Name      string    `gorm:not null`
	Email     string    `gorm:not null; unique`
	Gender    string    `gorm:not null`
	Team      string    `gorm:null`
	CPF       string    `gorm:unique;not null`
	StartDate time.Time `gorm:not null`
	BirthDate time.Time `gorm:not null`
	CreatedAt time.Time
	UpdatedAt time.Time
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

	err := e.validate()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Employee) validate() error {
	switch {
	case e.Name == "":
		return ErrInvalidName
	case e.Email == "":
		return ErrInvalidEmail
	case len(e.CPF) != 11:
		return ErrInvalidCPF
	case e.Gender == "":
		return ErrInvalidGender
	default:
		return nil
	}
}
