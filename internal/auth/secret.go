package auth

import (
	"encoding/json"
	"flag"
	"io/fs"
	"os"
)

var (
	SecretsFile = flag.String("secrets", "secrets.json", "Path to secrets.json.")
)

type Secret struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
}

func LoadSecrets() (map[string]Secret, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return testableLoadSecrets(os.DirFS(wd))
}

func testableLoadSecrets(fsys fs.FS) (map[string]Secret, error) {
	secrets := make(map[string]Secret)
	data, err := fs.ReadFile(fsys, *SecretsFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		return nil, err
	}
	return secrets, nil
}
