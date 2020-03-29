package middleware

import (
	"context"
	"github.com/gorilla/mux"
	userreg "github.com/wilsontwm/user-registration"
	"net/http"
	"strings"
	"user-registration-api/utils"
)

// Authenticate the authorization token in header
var JwtAuthentication = func() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			resp := make(map[string]interface{})
			// Check for authentication
			tokenHeader := r.Header.Get("Authorization")

			// If token is missing, then return error code 403 Unauthorized
			if tokenHeader == "" {
				utils.Fail(w, http.StatusUnauthorized, resp, "Missing authorization token.")
				return
			}

			// Check if the token format is correct, ie. Bearer {token}
			splitted := strings.Split(tokenHeader, " ")
			if len(splitted) != 2 {
				utils.Fail(w, http.StatusUnauthorized, resp, "Invalid authorization token format.")
				return
			}

			tokenPart := splitted[1] // Grab the second part
			token, err := userreg.Authenticate(tokenPart)

			if err != nil {
				utils.Fail(w, http.StatusUnauthorized, resp, err.Error())
				return
			}

			// Set the user ID in the context
			ctx := context.Background()
			ctx = context.WithValue(ctx, "userID", token.UserID)
			ctx = context.WithValue(ctx, "name", token.Name)
			ctx = context.WithValue(ctx, "email", token.Email)
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
	}
}
