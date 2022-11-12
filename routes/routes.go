package routes

import (
	"encoding/json"
	"gigmile-task/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func CountryRouter(r *mux.Router, app handler.HandlerModel) *mux.Router {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{"Server is running from docker compose"})
	}).Methods("GET")
	r.HandleFunc("/countries", app.CreateCountry).Methods("POST")
	r.HandleFunc("/countries", app.GetCountry).Methods("GET")
	r.HandleFunc("/countries/{id}", app.GetCountryById).Methods("GET")
	r.HandleFunc("/countries/{id}", app.UpdateCountry).Methods("PATCH")
	r.HandleFunc("/countries/{id}", app.DeleteCountryById).Methods("DELETE")
	return r

}
