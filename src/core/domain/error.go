package domain

import "fmt"

var (
	ErrInvalidCustomShortID = fmt.Errorf("invalid custom short ID")
	ErrInvalidURL           = fmt.Errorf("empty or invalid URL format")
	ErrInvalidInput         = fmt.Errorf("invalid input")
	ErrNotFound             = fmt.Errorf("resource not found")
	ErrDuplicateEntry       = fmt.Errorf("duplicate entry")
)
