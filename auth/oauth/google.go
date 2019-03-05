package oauth

import (
	"fmt"
	"net/http"
)

const (
	googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

func RedirectGoogle(w http.ResponseWriter, r *http.Request) {
	url := googleConfig.AuthCodeURL(oauthString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func CallbackGoogle(w http.ResponseWriter, r *http.Request) {
	userInfo, err := getUserInfo(r.FormValue("state"), r.FormValue("code"), googleUserInfoURL, googleConfig)
	fmt.Println(err)
	fmt.Println(userInfo)
	http.Redirect(w, r, "/auth", http.StatusTemporaryRedirect)
	// w.Write([]byte(userInfo))
}
