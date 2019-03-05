package auth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	googleClientID     = "GOOGLE_CLIENT_ID"
	googleClientSecret = "GOOGLE_CLIENT_SECRET"
)

var googleConfig = &oauth2.Config{
	ClientID:     os.Getenv(googleClientID),
	ClientSecret: os.Getenv(googleClientSecret),
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func redirectGoogle(w http.ResponseWriter, r *http.Request) {
	url := googleConfig.AuthCodeURL(oauthString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackGoogle(w http.ResponseWriter, r *http.Request) {
	userInfo, err := getGoogleUserInfo(r.FormValue("state"), r.FormValue("code"))
	fmt.Println(err)
	fmt.Println(userInfo)
	http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
	// w.Write([]byte(userInfo))
}

func getGoogleUserInfo(state string, code string) (string, error) {
	if state != oauthString {
		return "", fmt.Errorf("oauth - invalid state")
	}

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", fmt.Errorf("oauth - code exchange failed: %s", err.Error())
	}

	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
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
