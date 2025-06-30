package domain

import "fmt"

var (
	ErrInvalidCustomShortID = fmt.Errorf("invalid custom short ID")
	ErrInvalidURL           = fmt.Errorf("empty or invalid URL format")
)
