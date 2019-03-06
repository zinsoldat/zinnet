package oauth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/zinsoldat/zinnet-go/util"
)

func GetGoogleProvider() *OAuthProvider {
	return &OAuthProvider{
		Name: "google",
		config: &oauth2.Config{
			ClientID:     os.Getenv(googleClientID),
			ClientSecret: os.Getenv(googleClientSecret),
			RedirectURL:  "http://localhost:3000/auth/google/callback",
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
		UserInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
		state:       util.RandString(20),
	}
}
