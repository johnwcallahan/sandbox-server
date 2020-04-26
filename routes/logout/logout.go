package logout

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("DOMAIN")

	logoutURL, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutURL.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", os.Getenv("CLIENT_ID"))
	logoutURL.RawQuery = parameters.Encode()

	fmt.Println(logoutURL.String())

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}
