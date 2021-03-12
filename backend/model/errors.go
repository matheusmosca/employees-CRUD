package model

import "errors"

var (
	ErrInvalidName      = errors.New("The name field is required")
	ErrInvalidEmail     = errors.New("The email field is required")
	ErrInvalidCPF       = errors.New("The CPF field is required")
	ErrInvalidGender    = errors.New("The gender field is required")
	ErrInvalidBirthDate = errors.New("The birth date field is required")
	ErrInvalidStartDate = errors.New("The start date field is required")
)