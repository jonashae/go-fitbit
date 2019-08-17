package auth

import (
	"log"

	"golang.org/x/oauth2"
)

// PersistingTokenSource is a token source that persist token on refresh
type persistingTokenSource struct {
	original  oauth2.TokenSource
	persister TokenPersister
}

// Token returns a valid token
func (source persistingTokenSource) Token() (*oauth2.Token, error) {
	token, err := source.original.Token()
	if err != nil {
		return nil, err
	}

	err = source.persister.Persist(token)
	if err != nil {
		log.Printf("Error while persisting token: %s\n", err.Error())
	}

	return token, nil
}
