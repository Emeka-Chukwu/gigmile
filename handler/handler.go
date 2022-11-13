package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"gigmile-task/config"
	"gigmile-task/data"
	_ "gigmile-task/data"
	"gigmile-task/validator"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type HandlerModel struct {
	DB     *sql.DB
	Models data.Models
}
type jsonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (model *HandlerModel) CreateCountry(w http.ResponseWriter, r *http.Request) {
	config.Setheaders(w)
	var request data.Country

	app := config.Config{}
	err := json.NewDecoder(r.Body).Decode(&request)
	payLoad := jsonResponse{Success: false, Message: "Error parsing the entity", Data: err}
	if err != nil {
		app.WriteJSON(w, http.StatusBadRequest, payLoad)
		return
	}
	validate := validator.New()
	validate.Check(len(request.Name) > 1, "name", "must be atleast 2 characters")
	validate.Check(len(request.Continent) > 1, "continent", "must be atleast 2 characters")
	validate.Check(len(request.ShortName) > 1, "shortname", "must be atleast 2 characters")

	if !validate.Valid() {
		validate.FailedValidation(w, r)
		return
	}

	country, err := request.GetCountry(strings.ToLower(request.Name), strings.ToLower(request.ShortName))
	payLoad = jsonResponse{Success: false, Message: "Country already exist", Data: country}
	if err == nil {
		app.WriteJSON(w, http.StatusBadRequest, payLoad)
		return
	}
	resp, err := request.InsertCountry()

	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	payLoad = jsonResponse{Success: true, Message: "Country created successfully", Data: resp}
	app.WriteJSON(w, http.StatusOK, payLoad)
}

func (model *HandlerModel) GetCountry(w http.ResponseWriter, r *http.Request) {
	config.Setheaders(w)
	var request data.Country
	app := config.Config{}
	resp, err := request.GetAll()
	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	payLoad := jsonResponse{Success: true, Message: "Fetched countries successfully", Data: resp}
	app.WriteJSON(w, http.StatusOK, payLoad)
}
func (model *HandlerModel) GetDB() *sql.DB {
	return model.DB
}

func (model *HandlerModel) GetCountryById(w http.ResponseWriter, r *http.Request) {
	config.Setheaders(w)
	param := mux.Vars(r)
	id := param["id"]
	app := config.Config{}

	docId, err := strconv.Atoi(id)
	if err != nil {
		payLoad := jsonResponse{Success: false, Message: "invalid id", Data: err}

		app.WriteJSON(w, http.StatusBadRequest, payLoad)
		return

	}
	var request data.Country
	resp, err := request.GetOne(docId)
	if err != nil {
		app.ErrorJSON(w, errors.New("record not found"), http.StatusBadRequest)
		return
	}
	payLoad := jsonResponse{Success: true, Message: "Fetched countries successfully", Data: resp}
	app.WriteJSON(w, http.StatusOK, payLoad)
}

func (model *HandlerModel) UpdateCountry(w http.ResponseWriter, r *http.Request) {
	config.Setheaders(w)
	param := mux.Vars(r)
	id := param["id"]
	app := config.Config{}
	docId, err := strconv.Atoi(id)
	if err != nil {
		payLoad := jsonResponse{Success: false, Message: "invalid id", Data: err}

		app.WriteJSON(w, http.StatusBadRequest, payLoad)
		return

	}
	var request data.Country
	_ = json.NewDecoder(r.Body).Decode(&request)

	foundCountry, err := request.GetOne(docId)
	if err != nil {
		app.ErrorJSON(w, errors.New("record not found"), http.StatusBadRequest)
		return
	}
	// if &foundCountry.ID != &request.ID {
	// 	app.ErrorJSON(w, errors.New("record not found"), http.StatusBadRequest)
	// 	return
	// }
	foundCountry.Name = request.Name
	foundCountry.ShortName = request.ShortName
	foundCountry.Is_Operational = request.Is_Operational
	foundCountry.Continent = request.Continent
	foundCountry.Updated_At = time.Now()
	err = foundCountry.Update()

	if err != nil {
		app.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	payLoad := jsonResponse{Success: true, Message: "Country updated successfully", Data: foundCountry}
	app.WriteJSON(w, http.StatusOK, payLoad)
}

func (model *HandlerModel) DeleteCountryById(w http.ResponseWriter, r *http.Request) {
	config.Setheaders(w)
	param := mux.Vars(r)
	id := param["id"]
	app := config.Config{}

	docId, err := strconv.Atoi(id)
	if err != nil {
		payLoad := jsonResponse{Success: false, Message: "invalid id", Data: err}

		app.WriteJSON(w, http.StatusBadRequest, payLoad)
		return

	}
	var request data.Country
	err = request.DeleteByID(docId)
	if err != nil {
		app.ErrorJSON(w, errors.New("record not found"), http.StatusInternalServerError)
		return
	}
	payLoad := jsonResponse{Success: true, Message: "Country deleted successfully", Data: nil}
	app.WriteJSON(w, http.StatusOK, payLoad)
}
