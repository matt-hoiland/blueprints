package auth

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestLoadSecrets(t *testing.T) {
	tests := []struct {
		Name          string
		FS            fs.FS
		Secrets       map[string]Secret
		ErrorExpected bool
	}{
		{
			Name: "happy",
			FS: fstest.MapFS{
				"secrets.json": &fstest.MapFile{
					Data: []byte(`{"google":{"client_id":"abc","client_secret":"xyz","redirect_uri":"http://nowhere"}}`),
				},
			},
			Secrets: map[string]Secret{
				"google": {
					ClientID:     "abc",
					ClientSecret: "xyz",
					RedirectURI:  "http://nowhere",
				},
			},
		},
		{
			Name:          "missing file",
			FS:            fstest.MapFS{},
			ErrorExpected: true,
		},
		{
			Name: "malformed json",
			FS: fstest.MapFS{
				"secrets.json": &fstest.MapFile{
					Data: []byte(`{"google":{"client_id":"abc","client_secret":"xyz","redirect_uri":"http://nowhere}}`),
				},
			},
			ErrorExpected: true,
		},
		{
			Name: "empty json",
			FS: fstest.MapFS{
				"secrets.json": &fstest.MapFile{Data: []byte(`{}`)},
			},
			Secrets: make(map[string]Secret),
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			secrets, err := testableLoadSecrets(test.FS)
			if test.ErrorExpected {
				assert.Nil(t, secrets)
				assert.NotNil(t, err)
			} else {
				assert.Equal(t, test.Secrets, secrets)
				assert.Nil(t, err)
			}
		})
	}
}
