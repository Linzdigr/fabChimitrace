package middleware

import (
	"fab-chimix/application-go/conf"
	"net/http"
)

var notAuth = []string{ //List of endpoints that doesn't require auth
	conf.API_BASE_URL + "/user/create",
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		return
	})
}
