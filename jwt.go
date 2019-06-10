package main

import (
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
)

type jwtValidator struct {
	secret string
	idPURL *url.URL
}

func newJWTValidator(secret string, idPURL *url.URL) validator {
	return jwtValidator{
		secret: secret,
		idPURL: idPURL,
	}
}

func (v jwtValidator) authenticate(token Token, w http.ResponseWriter) bool {
	if token.token == "" {
		w.Header().Set("Location", v.idPURL.String())
		w.WriteHeader(http.StatusFound)
	}
	parsedToken, err := jwt.Parse(token.token, v.keyLookup)
	if err != nil {
		return false
	}

	if !parsedToken.Valid {
		return false
	}
	return false
}

func (v jwtValidator) keyLookup(_ *jwt.Token) (interface{}, error) {
	return []byte(v.secret), nil
}
