package middleware

import (
	"fmt"
	"net"
	"net/http"

	"github.com/erolkaldi/agency/pkg/models"
)

func AppKeyAuthorization(next http.Handler, apiConfig *models.Api) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		appKey := r.Header.Get("AppKey")
		if len(appKey) == 0 || !contains(apiConfig.AppKeys, appKey) {
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			fmt.Println(ip + ":" + r.RequestURI + " Invalid AppKey")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid AppKey"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
