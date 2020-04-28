package service

import "net/http"

/*
import (
	"context"
)
*/

type PrivilegeApiMiddleware func(PrivilegeApiService) PrivilegeApiService

/*
type logMiddleware struct {
	next UserApiService
}

func LogMiddleware() UserApiMiddleware {
	return func(next UserApiService) UserApiService {
		return &logMiddleware{next}
	}
}

func (l logMiddleware) HealthCheck() (string, error) {
	// Implement your middleware logic here

	return l.next.HealthCheck()
}
*/

//CORSMiddleware 支持跨域
func CORSMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, HEAD, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type,authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "false")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}
