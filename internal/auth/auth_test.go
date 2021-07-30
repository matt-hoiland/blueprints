package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matt-hoiland/blueprints/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		Name string
		Path string
		Code int
	}{
		{
			Name: "action login",
			Path: "/auth/login/provider",
			Code: http.StatusOK,
		},
		{
			Name: "action callback",
			Path: "/auth/callback/provider",
			Code: http.StatusNotFound,
		},
		{
			Name: "no provider",
			Path: "/auth/action",
			Code: http.StatusBadRequest,
		},
		{
			Name: "no action",
			Path: "/auth",
			Code: http.StatusBadRequest,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.Path, nil)
			w := httptest.NewRecorder()

			auth.LoginHandler(w, req)
			resp := w.Result()

			if !assert.NotNil(t, resp) {
				t.FailNow()
			}
			assert.Equal(t, test.Code, resp.StatusCode)
		})
	}
}
