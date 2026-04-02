package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rana-touseef11/go-chi-postgresql/internal/config"
)

type CustomClaims struct {
	// UserID string `json:"user_id"`
	// Email  string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(config.MustLoad().JWT_SECRET)

type contextKey string

const userKey contextKey = "jwt_user"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Get Authorization header
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// 2. Extract token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 3. Prepare claims
		claims := &CustomClaims{}

		// 4. Parse token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			// 🔒 Enforce HMAC signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("invalid signing method")
			}
			return jwtSecret, nil
		})

		// 5. Validate token
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// ✅ No need for manual exp check — jwt v5 handles it

		// 6. Store claims in context
		ctx := context.WithValue(r.Context(), userKey, claims)

		// 7. Call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JWTSign(userID string, exp time.Duration, roles ...string) (string, error) {
	var role string
	if len(roles) > 0 {
		role = roles[0]
	}

	claims := CustomClaims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app",
		},
	}
	// maps.Copy(claims, data)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	str := strings.Join([]string{"Bearer", jwtToken}, " ")
	return str, nil
}
