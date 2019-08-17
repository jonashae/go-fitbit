package fitbit

// Scope represents a specific resource permission
type Scope = string

// All available scopes
const (
	Activity  Scope = "activity"
	HeartRate Scope = "heartrate"
	Location  Scope = "location"
	Nutrition Scope = "nutrition"
	Profile   Scope = "profile"
	Settings  Scope = "settings"
	Sleep     Scope = "sleep"
	Social    Scope = "social"
	Weight    Scope = "weight"
)

// Config is the API client config
type Config struct {
	CallbackServerPort int
	ClientID           string
	ClientSecret       string
	Scopes             []Scope
	TokenPath          string
}
