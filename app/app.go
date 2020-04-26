package app

import (
	"encoding/gob"
	"os"

	"github.com/gorilla/sessions"
)

var (
	// Store is the global filesystem store.
	Store *sessions.FilesystemStore
)

// Init ...
func Init() error {
	Store = sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_SECRET")))
	gob.Register(map[string]interface{}{})
	return nil
}
