package unit_tests

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"middleware-jwt/core/authentication"
	"middleware-jwt/services"
	"middleware-jwt/services/models"
	"middleware-jwt/settings"
	"net/http"
	"os"
)

type AuthenticationServicesTestSuite struct{}

func (s *AuthenticationServicesTestSuite) SetUpSuite(c *C) {
	os.Setenv("GO_ENV", "tests")
	settings.Init()
}

func (suite *AuthenticationServicesTestSuite) TestLogin(c *C) {
	user := models.User{
		Username: "root",
		Password: "12345",
	}
	response, token := services.Login(&user)
	assert.Equal(t, http.StatusOK, response)
	assert.NotEmpty(t, token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectPassword(c *C) {
	user := models.User{
		Username: "root",
		Password: "Password",
	}
	response, token := services.Login(&user)
	assert.Equal(t, http.StatusUnauthorized, response)
	assert.Empty(t, token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginIncorrectUsername(c *C) {
	user := models.User{
		Username: "Username",
		Password: "12345",
	}
	response, token := services.Login(&user)
	assert.Equal(t, http.StatusUnauthorized, response)
	assert.Empty(t, token)
}

func (suite *AuthenticationServicesTestSuite) TestLoginEmptyCredentials(c *C) {
	user := models.User{
		Username: "",
		Password: "",
	}
	response, token := services.Login(&user)
	assert.Equal(t, http.StatusUnauthorized, response)
	assert.Empty(t, token)
}

func (suite *AuthenticationServicesTestSuite) TestRefreshToken(c *C) {
	user := &models.User{
		Username: "root",
		Password: "12345",
	}
	//authBackend := authentication.InitJWTAuthenticationBackend()
	//tokenString, err := authBackend.GenerateToken(user.UUID)
	//_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	return authBackend.PublicKey, nil
	//})
	//assert.Nil(t, err)

	newToken := services.RefreshToken(user)
	assert.NotEmpty(t, newToken)
}

func (suite *AuthenticationServicesTestSuite) TestRefreshTokenInvalidToken(c *C) {
	user := &models.User{
		Username: "root",
		Password: "12345",
	}
	//token := jwt.New(jwt.GetSigningMethod("RS256"))
	newToken := services.RefreshToken(user)
	assert.Empty(t, newToken)
}

func (suite *AuthenticationServicesTestSuite) TestLogout(c *C) {
	user := models.User{
		Username: "root",
		Password: "12345",
	}
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenString, err := authBackend.GenerateToken(user.UUID)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})

	err = authBackend.Logout(tokenString, token)
	assert.Nil(t, err)
}
