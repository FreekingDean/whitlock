package main

import (
	"net/http"
	"net/url"
	"os"
)

type validator interface {
	authenticate(Token, http.ResponseWriter) bool
}

type configuration struct {
	validator validator
}

func retreiveConfiguration() configuration {
	//http.get from whitlock.io
	jwtSecret := os.Getenv("JWT_SECRET")
	idPURL, _ := url.Parse(os.Getenv("IDP_URL"))

	return configuration{
		validator: newJWTValidator(jwtSecret, idPURL),
	}
}
