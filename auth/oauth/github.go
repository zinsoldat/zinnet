package oauth

import (
	"fmt"
	"net/http"
)

const (
	githubUserInfoURL = "https://api.github.com/user/emails"
)

func RedirectGithub(w http.ResponseWriter, r *http.Request) {
	url := githhubConfig.AuthCodeURL(oauthString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func CallbackGithub(w http.ResponseWriter, r *http.Request) {
	userInfo, _ := getUserInfo(r.FormValue("state"), r.FormValue("code"), githubUserInfoURL, githhubConfig)
	fmt.Println(userInfo)
	http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
	// w.Write([]byte(userInfo))
}
