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
	CompleteAuth(data objx.Map) (*common.Credentials, error)
	GetUser(*common.Credentials) (common.User, error)
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
	providerName := segs[3]
	switch action {
	case "login":
		provider, err := h.oauthAdapter.Provider(providerName)
		if err != nil {
			log.Fatalln("Error when trying to get provider", providerName, "-", err)
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("Error when trying to GetBeginAuthURL for", providerName, "-", err)
		}
		w.Header().Add("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		provider, err := h.oauthAdapter.Provider(providerName)
		if err != nil {
			log.Fatalln("Error when trying to get provider", providerName, "-", err)
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("Error when trying to complete auth for", providerName, "-", err)
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("Error when trying to get user from", providerName, "-", err)
		}

		authCookieValue := objx.New(map[string]interface{}{
			"name": user.Name(),
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})

		w.Header().Add("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}

type GomniAuthAdapter struct{}

func (*GomniAuthAdapter) Provider(provider string) (Provider, error) {
	return gomniauth.Provider(provider)
}
