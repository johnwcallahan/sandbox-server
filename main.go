package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/johnwcallahan/sandbox-server/app"
	"github.com/johnwcallahan/sandbox-server/routes/callback"
	"github.com/johnwcallahan/sandbox-server/routes/login"
	"github.com/johnwcallahan/sandbox-server/routes/logout"
	"github.com/johnwcallahan/sandbox-server/routes/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app.Init()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/callback", callback.Handler)
	r.HandleFunc("/login", login.Handler)
	r.HandleFunc("/user", user.Handler)
	r.HandleFunc("/logout", logout.Handler)

	fmt.Println("Listening on port " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}

// HomeHandler handles the root route.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")

}
