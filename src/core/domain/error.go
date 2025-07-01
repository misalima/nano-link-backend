package domain

import "fmt"

var (
	ErrInvalidCustomShortID = fmt.Errorf("invalid custom short ID")
	ErrInvalidURL           = fmt.Errorf("empty or invalid URL format")
	ErrInvalidInput         = fmt.Errorf("invalid input")
	ErrURLNotFound          = fmt.Errorf("URL not found")
	ErrNotFound             = fmt.Errorf("resource not found")
	ErrDuplicateEntry       = fmt.Errorf("duplicate entry")
	ErrTagNotFound          = fmt.Errorf("tag not found")
)
