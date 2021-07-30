package auth_test

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matt-hoiland/blueprints/internal/auth"
	"github.com/matt-hoiland/blueprints/internal/auth/mocks"
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	tests := []struct {
		Name            string
		Path            string
		Code            int
		SetExpectations func(adapter *mocks.MockOAuthAdapter, provider *mocks.MockProvider)
		CheckResponse   func(t *testing.T, resp *http.Response)
	}{
		{
			Name: "action login",
			Path: "/auth/login/provider",
			Code: http.StatusTemporaryRedirect,
			SetExpectations: func(adapter *mocks.MockOAuthAdapter, provider *mocks.MockProvider) {
				adapter.EXPECT().Provider("provider").Return(provider, nil).Times(1)
				provider.EXPECT().GetBeginAuthURL(nil, nil).Return("http://nowhere", nil).Times(1)
			},
			CheckResponse: func(t *testing.T, resp *http.Response) {
				loc, err := resp.Location()
				if !assert.Nil(t, err) {
					t.FailNow()
				}
				assert.Equal(t, "http://nowhere", loc.String())
			},
		},
		{
			Name: "action callback",
			Path: "/auth/callback/provider?code=12345",
			Code: http.StatusTemporaryRedirect,
			SetExpectations: func(adapter *mocks.MockOAuthAdapter, provider *mocks.MockProvider) {
				creds := &common.Credentials{Map: objx.Map{}}

				adapter.EXPECT().Provider("provider").Return(provider, nil).Times(1)
				provider.EXPECT().CompleteAuth(gomock.Any()).Return(creds, nil).Times(1)
				provider.EXPECT().GetUser(creds).Return(&mocks.MockUser{Username: "John Doe"}, nil).Times(1)
			},
			CheckResponse: func(t *testing.T, resp *http.Response) {
				if !assert.Len(t, resp.Cookies(), 1) {
					t.FailNow()
				}
				cookie := resp.Cookies()[0]
				assert.Equal(t, "auth", cookie.Name)
				assert.Equal(t, base64.StdEncoding.EncodeToString([]byte(`{"name":"John Doe"}`)), cookie.Value)
				assert.Equal(t, "/", cookie.Path)
				loc, err := resp.Location()
				if !assert.Nil(t, err) {
					t.FailNow()
				}
				assert.Equal(t, "/chat", loc.String())
			},
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
			c := gomock.NewController(t)
			defer c.Finish()

			req := httptest.NewRequest("GET", test.Path, nil)
			w := httptest.NewRecorder()

			adapter := mocks.NewMockOAuthAdapter(c)
			provider := mocks.NewMockProvider(c)
			if test.SetExpectations != nil {
				test.SetExpectations(adapter, provider)
			}

			handler := auth.NewLoginHandler(adapter)
			handler.ServeHTTP(w, req)
			resp := w.Result()

			if !assert.NotNil(t, resp) {
				t.FailNow()
			}
			assert.Equal(t, test.Code, resp.StatusCode)
			if test.CheckResponse != nil {
				test.CheckResponse(t, resp)
			}
		})
	}
}
