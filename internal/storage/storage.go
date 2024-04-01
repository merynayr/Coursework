package storage

import "errors"

var (
	ErrNotFound = errors.New("object not found")
	ErrExists   = errors.New("ogject exists")
)
