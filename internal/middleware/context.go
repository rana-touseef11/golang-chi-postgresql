package middleware

import (
	"net/http"

	"github.com/rana-touseef11/go-chi-postgresql/pkg/validator"
)

func GetBody[T any](r *http.Request) (*T, bool) {
	u, ok := r.Context().Value(validator.BodyKey).(T)
	return &u, ok
}
