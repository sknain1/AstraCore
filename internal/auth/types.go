package auth

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Code         int    `json:"code"`
	Message      string `json:"message"`
	Status       string `json:"s"`
}
