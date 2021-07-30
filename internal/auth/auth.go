package auth

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

type OAuthAdapter interface {
	Provider(string) (Provider, error)
}

type Provider interface {
	GetBeginAuthURL(*common.State, objx.Map) (string, error)
}

type loginHandler struct {
	oauthAdapter OAuthAdapter
}

func NewLoginHandler(oauthAdapter OAuthAdapter) http.Handler {
	return &loginHandler{oauthAdapter: oauthAdapter}
}

// LoginHandler handles the third-party login process
// format: /auth/{action}/{provider}
func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Expected path patter: /auth/{action}/{provider}")
		return
	}
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

type GomniAuthAdapter struct{}

func (*GomniAuthAdapter) Provider(provider string) (Provider, error) {
	return gomniauth.Provider(provider)
}
