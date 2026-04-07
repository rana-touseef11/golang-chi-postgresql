package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response[T any] struct {
	Status  string   `json:"status"`
	Message string   `json:"message,omitempty"`
	Data    T        `json:"data,omitempty"`
	Errors  []string `json:"errors,omitempty"`
	Meta    any      `json:"meta,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response[any] {
	return Response[any]{
		Status: StatusError,
		Errors: []string{err.Error()},
	}
}

func ValidationErrors(errs validator.ValidationErrors) Response[any] {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}

	return Response[any]{
		Status:  StatusError,
		Errors:  errMsgs,
		Message: "Validation Faild",
	}
}
