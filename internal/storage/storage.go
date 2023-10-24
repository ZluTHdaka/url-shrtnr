package storage

import "errors"

var (
	ErrURLNotFound = errors.New("Url not found")
	ErrURLExists   = errors.New("Url aalready exists")
)
