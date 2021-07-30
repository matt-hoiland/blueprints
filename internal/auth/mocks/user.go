// Hand crafted mock of common.User
package mocks

import (
	"github.com/stretchr/gomniauth/common"
	"github.com/stretchr/objx"
)

type MockUser struct {
	Username string
}

func (m *MockUser) AuthCode() string  { return "" }
func (m *MockUser) Email() string     { return "" }
func (m *MockUser) Name() string      { return m.Username }
func (m *MockUser) Nickname() string  { return "" }
func (m *MockUser) AvatarURL() string { return "" }
func (m *MockUser) ProviderCredentials() map[string]*common.Credentials {
	return make(map[string]*common.Credentials)
}
func (m *MockUser) IDForProvider(provider string) string { return "" }
func (m *MockUser) Data() objx.Map                       { return make(objx.Map) }
