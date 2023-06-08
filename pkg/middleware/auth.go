package middleware

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/erolkaldi/agency/pkg/service"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			fmt.Println(ip + ":" + r.RequestURI + " Missing Authorization Header")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := service.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(ip + ":" + r.RequestURI + " Error verifying JWT token: " + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		expirationTime := time.Unix(int64(claims.(jwt.MapClaims)["expiration"].(float64)), 0)
		if time.Now().After(expirationTime) {
			fmt.Println(ip + ":" + r.RequestURI + " Token expired")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token expired"))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		ID := claims.(jwt.MapClaims)["id"].(string)
		email := claims.(jwt.MapClaims)["email"].(string)

		r.Header.Set("name", name)
		r.Header.Set("userID", ID)
		r.Header.Set("email", email)
		fmt.Println(ip + ":" + r.RequestURI + " Token Accepted")
		next.ServeHTTP(w, r)
	})
}
