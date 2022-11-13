package validator

import (
	"gigmile-task/config"
	"net/http"
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validator) FailedValidation(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}
	payload.Success = false
	payload.Message = "Validation failed"
	payload.Errors = v.Errors

	app := config.Config{}
	app.WriteJSON(w, http.StatusUnprocessableEntity, payload)
}
