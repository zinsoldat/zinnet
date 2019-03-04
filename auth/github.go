package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	githubClientID     = "GITHUB_CLIENT_ID"
	githubClientSecret = "GITHUB_CLIENT_SECRET"
)

var githhubConfig = &oauth2.Config{
	ClientID:     os.Getenv(githubClientID),
	ClientSecret: os.Getenv(githubClientSecret),
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}

var oauthString = "random-string"

func redirectGithub(w http.ResponseWriter, r *http.Request) {
	url := githhubConfig.AuthCodeURL(oauthString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackGithub(w http.ResponseWriter, r *http.Request) {
	userInfo, _ := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	fmt.Println(userInfo)
	http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
	// w.Write([]byte(userInfo))
}

func getUserInfo(state string, code string) (string, error) {
	if state != oauthString {
		return "", fmt.Errorf("oauth - invalid state")
	}

	token, err := githhubConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("oauth - code exchange failed: %s", err.Error())
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user/emails", nil)
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
