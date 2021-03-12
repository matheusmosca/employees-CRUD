package model

import "errors"

var (
	ErrInvalidName      = errors.New("The name field is required")
	ErrInvalidEmail     = errors.New("The email field is required")
	ErrInvalidCPF       = errors.New("The CPF field must have 11 numbers")
	ErrInvalidGender    = errors.New("The gender field is required")
	ErrInvalidTeam      = errors.New("Provide a valid team value")
	ErrInvalidBirthDate = errors.New("The birth date field is required")
	ErrInvalidStartDate = errors.New("The start date field is required")
)
