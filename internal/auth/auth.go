package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

const timeout = 30 * time.Second

// ImplicitGrantFlow is recommended for apps without web service
var ImplicitGrantFlow = oauth2.SetAuthURLParam("response_type", "token")

// Handler handles the authn and authz with Fitbit API
type Handler struct {
	config       *oauth2.Config
	callbackPort int
	storage      TokenStorage
}

// NewHandler creates new auth handler
func NewHandler(client string, secret string, scopes []string, callbackPort int, storage TokenStorage) *Handler {
	config := &oauth2.Config{
		ClientID:     client,
		ClientSecret: secret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.fitbit.com/oauth2/authorize",
			TokenURL: "https://api.fitbit.com/oauth2/token",
		},
	}
	return &Handler{config: config, storage: storage, callbackPort: callbackPort}
}

// Client returns an authenticated http client
func (handler Handler) Client(context context.Context) (*http.Client, error) {
	token, err := handler.acquireToken()
	if err != nil {
		return nil, err
	}

	tokenSource := handler.createTokenSource(context, token)

	return oauth2.NewClient(context, tokenSource), nil
}

func (handler *Handler) acquireToken() (*oauth2.Token, error) {
	token, err := handler.loadStoredToken()
	if err != nil {
		return nil, err
	}

	if token != nil {
		return token, nil
	}

	return handler.requestInitialToken()
}

func (handler *Handler) loadStoredToken() (*oauth2.Token, error) {
	if handler.storage == nil {
		return nil, nil
	}

	token, err := handler.storage.Load()
	if err != nil {
		log.Printf("Error loading token from storage: %s\n", err.Error())
	}

	return token, err
}

func (handler *Handler) requestInitialToken() (*oauth2.Token, error) {
	state, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	authURL := handler.config.AuthCodeURL(state.String())
	log.Printf("Visit URL: %s\n", authURL)

	callback := NewCallbackServer(handler.callbackPort)
	context, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	code, err := callback.AwaitCode(context)
	if err != nil {
		return nil, err
	}

	return handler.config.Exchange(context, code, ImplicitGrantFlow)
}

func (handler *Handler) createTokenSource(context context.Context, token *oauth2.Token) oauth2.TokenSource {
	original := handler.config.TokenSource(context, token)
	return persistingTokenSource{original: original, persister: handler.storage}
}
