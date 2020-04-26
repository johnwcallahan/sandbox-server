package user

import (
	"net/http"

	"github.com/johnwcallahan/sandbox-server/app"
	"github.com/johnwcallahan/sandbox-server/routes/templates"
)

// Handler ...
func Handler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "user", session.Values["profile"])
}
