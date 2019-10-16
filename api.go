package scomportal

// API ..
type API struct {
	Email   string
	Token   string
	BaseURL string
}

// NewAPI ..
func NewAPI(email, token string) *API {
	return &API{
		Email:   email,
		Token:   token,
		BaseURL: "https://portal.servers.com/rest",
	}
}
