package middleware

import "net/http"

//Cors config enabled to all origin
func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, PUT, GET")
		response.Header().Set("Content-Type", "application/json")
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if request.Method == "OPTIONS" {
			response.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(response, request)
		}
	})
}
