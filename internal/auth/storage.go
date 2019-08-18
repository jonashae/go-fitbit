package auth

import (
	"golang.org/x/oauth2"
)

// TokenPersister is an interface for persisting tokens
type TokenPersister interface {
	Persist(*oauth2.Token) error
}

// TokenLoader is an interface for loading tokens
type TokenLoader interface {
	Load() (*oauth2.Token, error)
}

// TokenCleaner is an interface for deleting tokens
type TokenCleaner interface {
	Wipe() error
}

// TokenStorage is a union of TokenPersister and TokenLoader
type TokenStorage interface {
	TokenCleaner
	TokenLoader
	TokenPersister
}
