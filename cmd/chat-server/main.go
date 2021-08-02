package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/matt-hoiland/blueprints/internal/auth"
	"github.com/matt-hoiland/blueprints/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

var (
	host    = flag.String("host", "localhost", "The host of the chat server.")
	port    = flag.String("port", "8080", "The port of the chat server.")
	verbose = flag.Bool("v", false, "Enable verbose logging output")
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		var err error
		t.templ, err = template.ParseFiles(filepath.Join("templates", t.filename))
		if err != nil {
			panic(err)
		}
	})
	err := t.templ.Execute(w, r)
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	secrets, err := auth.LoadSecrets()
	if err != nil {
		panic(err)
	}

	gomniauth.SetSecurityKey("TODO replace with secure passphrase or hash")
	gomniauth.WithProviders(
		google.New(secrets["google"].ClientID, secrets["google"].ClientSecret, secrets["google"].RedirectURI),
		github.New(secrets["github"].ClientID, secrets["github"].ClientSecret, secrets["github"].RedirectURI),
	)

	addr := fmt.Sprintf("%s:%s", *host, *port)
	wd, _ := os.Getwd()
	log.Println("Working Directory", wd)

	r := newRoom()
	if *verbose {
		r.tracer = trace.New(os.Stdout)
	}
	http.Handle("/chat", auth.MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/auth/", auth.NewLoginHandler(&auth.GomniAuthAdapter{}))
	http.Handle("/room", r)
	// Get the room going
	go r.run()
	// Start the web server
	log.Println("Starting web server on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
