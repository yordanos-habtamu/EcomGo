package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/yordanos-habtamu/EcomGo.git/config"
	"github.com/yordanos-habtamu/EcomGo.git/utils"
)

var (
	// Secret key for sessions, should be stored in an environment variable for production.
	secretKey = []byte(config.Envs.JWT_SECRET)
	store     = sessions.NewCookieStore(secretKey)
)

type CustomClaims struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// GetJWTFromSession retrieves the JWT token from the session
func GetJWTFromSession(r *http.Request) (string, error) {
	session, err := store.Get(r, "EcomGo")
	if err != nil {
		return "", err
	}

	// Retrieve the token from the session
	tokenString, ok := session.Values["token"].(string)
	if !ok {
		return "", nil
	}
	return tokenString, nil
}

// SetJWTInSession stores the JWT token in the session
func SetJWTInSession(w http.ResponseWriter, r *http.Request, tokenString string) error {
	session, err := store.Get(r, "EcomGo")
	if err != nil {
		return err
	}

	// Store the JWT token in the session
	session.Values["token"] = tokenString
	session.Save(r, w)
	return nil
}

// JwtMiddleware is a middleware to check for valid JWT tokens and enforce RBAC
func JwtMiddleware(requiredRole string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("authorization header is missing"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 {
			// The token is the second part of the split result
			tokenString := parts[1]
			token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil {
				utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("not authorized"))
				return
			}

			// Extract claims and check if the token is valid
			claims, ok := token.Claims.(*CustomClaims)
			if !ok || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Add the user information to the context
			ctx := context.WithValue(r.Context(), "user", claims)
			r = r.WithContext(ctx)

			// Check the user's role if required
			if requiredRole != "" && claims.Role != requiredRole {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Proceed to the next handler
			// You need to make sure the next handler is passed as an argument to this middleware
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Your actual handler logic goes here
			})
			next.ServeHTTP(w, r)
			return
		} else {
			// Handle error if the header is not in the correct format
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization header"))
			return
		}
	}
}
