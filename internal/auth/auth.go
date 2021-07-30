package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// LoginHandler handles the third-party login process
// format: /auth/{action}/{provider}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		log.Println("TODO handle login for", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
