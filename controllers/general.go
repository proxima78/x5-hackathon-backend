package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/proxima78/x5-hackathon-backend/app"
	httpModels "github.com/proxima78/x5-hackathon-backend/models/http"
	"github.com/proxima78/x5-hackathon-backend/services"
)

type GeneralController struct {
	app            *app.App
	generalService *services.GeneralService
}

func NewGeneralController(a *app.App, gs *services.GeneralService) *GeneralController {
	return &GeneralController{
		app:            a,
		generalService: gs,
	}
}

func (gc *GeneralController) CheckAccountExists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var checkAccountExistsRequest httpModels.CheckAccountExistsRequest
	err := decoder.Decode(&checkAccountExistsRequest)
	if err != nil {
		log.Println(err)
		ReturnHTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := gc.generalService.CheckAccountExists(&checkAccountExistsRequest)
	if err != nil {
		log.Println(err)
		ReturnHTTPError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, _ := json.Marshal(res)
	w.Write(json)
}

func (gc *GeneralController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var createAccountRequest httpModels.CreateAccountRequest
	err := decoder.Decode(&createAccountRequest)
	if err != nil {
		log.Println(err)
		ReturnHTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := gc.generalService.CreateAccount(createAccountRequest)
	if err != nil {
		log.Println(err)
		ReturnHTTPError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json, _ := json.Marshal(res)
	w.Write(json)
}
