package fitbit

// Scope represents a specific resource permission
type Scope = string



// Config is the API client config
type Config struct {
	CallbackServerPort int
	ClientID           string
	ClientSecret       string
	Scopes             []string
	TokenPath          string
}
