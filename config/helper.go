package config

import (
	"encoding/json"
	"errors"

	"io"
	"net/http"
)

type jsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxByte := 1048578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(data)
	if err != io.EOF {
		return errors.New("body must have only a single JSON value")
	}
	return nil
}

func (app *Config) WriteJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

func (app *Config) ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload jsonResponse
	payload.Success = false
	payload.Data = err.Error()
	return app.WriteJSON(w, statusCode, payload)
}

func Setheaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE")
	w.Header().Set("Content-Type", "application/json")
}
