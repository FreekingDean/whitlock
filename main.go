package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var targetHost = os.Getenv("PROXY_TARGET")

func main() {
	port := os.Getenv("PORT")
	rev := &httputil.ReverseProxy{
		Director: reverser,
	}
	log.Fatal(http.ListenAndServe(":"+port, auther(rev)))
}

func auther(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := parseToken(r)
		if validToken(token) {
			next.ServeHTTP(w, r)
		}
	})
}

type Token struct {
	style string
	token string
}

func validToken(token Token) bool {
	return true
}

func parseToken(r *http.Request) Token {
	headerToken := strings.Split(r.Header.Get("Authorization"), "=")
	if len(headerToken) != 2 {
		return Token{}
	}
	return Token{
		style: headerToken[0],
		token: headerToken[1],
	}
}

func reverser(req *http.Request) {
	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}
	req.URL.Host = targetHost
	req.Host = targetHost
	if _, ok := req.Header["User-Agent"]; !ok {
		req.Header.Set("User-Agent", "")
	}
}
