package handler_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"gigmile-task/data"
	"gigmile-task/handler"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountryCreate(t *testing.T) {
	t.Run("Populate Db with country values", func(t *testing.T) {

		// con := config.ConnectToDB()

		hand := handler.HandlerModel{}
		clearTable(hand.DB)

		countries := []data.Country{
			{
				Name:           "Nigeria",
				ShortName:      "NG",
				Is_Operational: true,
				Continent:      "africa",
			},
		}

		for _, country := range countries {
			payload, _ := json.Marshal(country)
			url := "http://localhost:8080/countries"
			req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			assert.Equal(t, nil, err)
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		}
		for _, country := range countries {
			payload, _ := json.Marshal(country)
			url := "http://localhost:8080/countries"
			req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			assert.Equal(t, nil, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		}
	})

}

func TestCountryReadAll(t *testing.T) {
	t.Run("Populate Db with country values", func(t *testing.T) {

		url := "http://localhost:8080/countries"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Equal(t, nil, err)
		// assert.Equal(t, countries[0].Name, res)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})

}

func TestCountryReadOne(t *testing.T) {
	t.Run("Populate Db with country values", func(t *testing.T) {

		url := "http://localhost:8080/countries/90"
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

	})

}

func clearTable(a *sql.DB) {
	a.Exec("DELETE FROM countries")
	a.Exec("ALTER SEQUENCE countries_id_seq RESTART WITH 1")
}
