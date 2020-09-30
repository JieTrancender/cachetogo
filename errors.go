package cachetogo

import (
	"errors"
)

var (
	// ErrKeyNotFound key is not found in cache
	ErrKeyNotFound = errors.New("Key not found in cache")

	// ErrkeyNotFoundOrLoadable key is not found in cache and could not be loaded itno cache
	ErrkeyNotFoundOrLoadable = errors.New("Key not found and could not be loaded into cache")
)
