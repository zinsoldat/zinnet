package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

const (
	googleClientID     = "GOOGLE_CLIENT_ID"
	googleClientSecret = "GOOGLE_CLIENT_SECRET"
	githubClientID     = "GITHUB_CLIENT_ID"
	githubClientSecret = "GITHUB_CLIENT_SECRET"
	oauthString        = "random-string"
)

var (
	githhubConfig = &oauth2.Config{
		ClientID:     os.Getenv(githubClientID),
		ClientSecret: os.Getenv(githubClientSecret),
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}

	googleConfig = &oauth2.Config{
		ClientID:     os.Getenv(googleClientID),
		ClientSecret: os.Getenv(googleClientSecret),
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func getUserInfo(state string, code string, url string, config *oauth2.Config) (string, error) {
	if state != oauthString {
		return "", fmt.Errorf("oauth - invalid state")
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("oauth - code exchange failed: %s", err.Error())
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	token.SetAuthHeader(req)
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("oauth - failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("oauth - failed to read response body: %s", err.Error())
	}
	return string(contents), nil
}
