package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

const (
	googleClientID     = "GOOGLE_CLIENT_ID"
	googleClientSecret = "GOOGLE_CLIENT_SECRET"
	githubClientID     = "GITHUB_CLIENT_ID"
	githubClientSecret = "GITHUB_CLIENT_SECRET"
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

type OAuthProvider struct {
	Name            string
	config          *oauth2.Config
	UserInfoURL     string
	state           string
	parseUserInfo   func(userInfo string) (string, error)
	RedirectHandler func(c echo.Context) error
	CallbackHandler func(c echo.Context) error
}

func (provider *OAuthProvider) getUserInfo(state string, code string) (string, error) {

	if state != provider.state {
		return "", fmt.Errorf("oauth - invalid state")
	}

	token, err := provider.config.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("oauth - code exchange failed: %s", err.Error())
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", provider.UserInfoURL, nil)
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

type OAuthCallbackData struct {
	state string
	code  string
}

func (provider *OAuthProvider) Callback(c echo.Context) error {
	data := new(OAuthCallbackData)
	if err := c.Bind(data); err != nil {
		return err
	}
	userInfo, _ := provider.getUserInfo(data.state, data.code)
	fmt.Println(userInfo)
	c.Redirect(http.StatusTemporaryRedirect, "/auth")
	return nil
}

func (provider *OAuthProvider) Redirect(c echo.Context) error {
	url := provider.config.AuthCodeURL(provider.state)
	c.Redirect(http.StatusTemporaryRedirect, url)
	return nil
}
