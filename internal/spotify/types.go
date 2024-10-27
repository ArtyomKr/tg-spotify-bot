package spotify

type Client struct {
	clientID     string
	clientSecret string
	redirectURI  string
	apiURL       string
}

type TokenResBody struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
