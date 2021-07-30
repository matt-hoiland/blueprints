package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matt-hoiland/blueprints/internal/auth"
	"github.com/stretchr/testify/assert"
)

type MockHTTPHandler struct {
	Called bool
}

func (mock *MockHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mock.Called = true
}

func TestAuthHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		Name          string
		BuildRequest  func() *http.Request
		CheckResponse func(t *testing.T, resp *http.Response)
		MockCalled    bool
	}{
		{
			Name: "redirect",
			BuildRequest: func() *http.Request {
				return httptest.NewRequest("GET", "/chat", nil)
			},
			CheckResponse: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusTemporaryRedirect, resp.StatusCode)
				path, err := resp.Location()
				if !assert.Nil(t, err) {
					t.Error(err)
					t.FailNow()
				}
				assert.Contains(t, path.Path, "/login")
			},
		},
		{
			Name: "passthrough",
			BuildRequest: func() *http.Request {
				req := httptest.NewRequest("GET", "/chat", nil)
				req.AddCookie(&http.Cookie{Name: "auth", Value: "nonce"})
				return req
			},
			CheckResponse: func(t *testing.T, resp *http.Response) {
				assert.Equal(t, http.StatusOK, resp.StatusCode)
			},
			MockCalled: true,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.Name, func(t *testing.T) {
			next := &MockHTTPHandler{}
			req := test.BuildRequest()
			w := httptest.NewRecorder()

			handler := auth.MustAuth(next)
			handler.ServeHTTP(w, req)

			resp := w.Result()
			test.CheckResponse(t, resp)
			assert.Equal(t, next.Called, test.MockCalled)
		})
	}
}
