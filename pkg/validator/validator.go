package validator

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rana-touseef11/go-chi-postgresql/internal/response"
)

type contextKey string

const BodyKey contextKey = "body"

func ValidateRequest[T any]() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var body T

			// Decode Json
			err := json.NewDecoder(r.Body).Decode(&body)
			if errors.Is(err, io.EOF) {
				http.Error(w, "Empty Body", http.StatusBadRequest)
				return
			}
			if err != nil {
				http.Error(w, "Invaild Request", http.StatusBadRequest)
				return
			}

			if err := validator.New().Struct(body); err != nil {
				validationErr := err.(validator.ValidationErrors)
				resErr := response.ValidationErrors(validationErr)
				response.WriteJson(w, http.StatusBadRequest, resErr)
				return
			}

			ctx := context.WithValue(r.Context(), BodyKey, body)

			// continue
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
