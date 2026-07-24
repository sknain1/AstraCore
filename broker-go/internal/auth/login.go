package auth

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	appID := os.Getenv("FYERS_APP_ID")
	redirectURI := url.QueryEscape(os.Getenv("FYERS_REDIRECT_URI"))

	loginURL := fmt.Sprintf(
		"https://api-t1.fyers.in/api/v3/generate-authcode?client_id=%s&redirect_uri=%s&response_type=code&state=AstraCore",
		appID,
		redirectURI,
	)

	fmt.Println(loginURL)

	http.Redirect(w, r, loginURL, http.StatusFound)
}
