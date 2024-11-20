package f

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func WriteToJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func WriteToJsonError(w http.ResponseWriter, r *http.Request, data interface{}) {
	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
}

type ErrorType struct {
	Field     string
	Condition string
}

func ErrorValidation(err error) []ErrorType {
	var errors []ErrorType
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			errors = append(errors, ErrorType{
				Field:     e.Field(),
				Condition: e.Tag(),
			})
		}
	}
	return errors
}
