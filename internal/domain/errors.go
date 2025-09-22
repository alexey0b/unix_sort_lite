package domain

import "errors"

var (
	ErrConflictOpts  = errors.New("sort: conflicting sort options")
	ErrWrongOrder    = errors.New("sort: wrong order")
	ErrInvalideField = errors.New("sort: invalid number of field")
)
