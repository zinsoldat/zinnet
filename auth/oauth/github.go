package oauth

import (
	"os"

	"github.com/zinsoldat/zinnet-go/util"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func GetGithubProvider() *OAuthProvider {
	return &OAuthProvider{
		Name: "github",
		config: &oauth2.Config{
			ClientID:     os.Getenv(githubClientID),
			ClientSecret: os.Getenv(githubClientSecret),
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
		UserInfoURL: "https://api.github.com/user/emails",
		state:       util.RandString(20),
	}
}
