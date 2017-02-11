package unit_tests

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"middleware-jwt/core/authentication"
	"middleware-jwt/core/redis"
	"middleware-jwt/services/models"
	"middleware-jwt/settings"
	"os"
	"testing"
)

type AuthenticationBackendTestSuite struct{}

var t *testing.T

func (s *AuthenticationBackendTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *AuthenticationBackendTestSuite) TestInitJWTAuthenticationBackend(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	c.Assert(authBackend, NotNil)
	c.Assert(authBackend.PublicKey, NotNil)
}

func (suite *AuthenticationBackendTestSuite) TestGenerateToken(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken("1234")

	assert.Nil(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})

	assert.Nil(t, err)
	assert.True(t, token.Valid)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticate(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := &models.User{
		Username: "root",
		Password: "12345",
	}
	c.Assert(authBackend.Authenticate(user), Equals, true)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectPass(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := models.User{
		Username: "root",
		Password: "Password",
	}
	c.Assert(authBackend.Authenticate(&user), Equals, false)
}

func (suite *AuthenticationBackendTestSuite) TestAuthenticateIncorrectUsername(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	user := &models.User{
		Username: "Username",
		Password: "12345",
	}
	c.Assert(authBackend.Authenticate(user), Equals, false)
}

func (suite *AuthenticationBackendTestSuite) TestLogout(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken("1234")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)
	assert.Nil(t, err)

	redisConn := redis.Connect()
	redisValue, err := redisConn.GetValue(tokenString)
	assert.Nil(t, err)
	assert.NotEmpty(t, redisValue)
}

func (suite *AuthenticationBackendTestSuite) TestIsInBlacklist(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken("1234")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	err = authBackend.Logout(tokenString, token)
	assert.Nil(t, err)

	assert.True(t, authBackend.IsInBlacklist(tokenString))
}

func (suite *AuthenticationBackendTestSuite) TestIsNotInBlacklist(c *C) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	assert.False(t, authBackend.IsInBlacklist("1234"))
}
