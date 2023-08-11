package path

// Token holds the data returned from the /token endpoint after successful authentication
type Token struct {
	AccessToken string                 `json:"access_token"`
	TokenType   AuthorizationTokenType `json:"token_type"`
}

// AccessTokenRequest holds the necessary data that the /token endpoint expects
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`

	// Required
	Username string `json:"username"`
	Password string `json:"password"`

	Scope        string `json:"scope"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
