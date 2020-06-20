package router

import (
	"github.com/proxima78/x5-hackathon-backend/controllers"
	"github.com/proxima78/x5-hackathon-backend/repositories"
	"github.com/proxima78/x5-hackathon-backend/services"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/proxima78/x5-hackathon-backend/app"
	"github.com/proxima78/x5-hackathon-backend/middlewares"
)

func NewRouter(a *app.App) (*mux.Router, error) {
	r := mux.NewRouter()

	// NOTE Create repositories here
	ur, err := repositories.NewUserRepository(a)
	if err != nil {
		return nil, err
	}

	// NOTE Create services here
	gs := services.NewGeneralService(a, ur)

	// NOTE Create controllers here
	gc := controllers.NewGeneralController(a, gs)

	// NOTE Create middlewares here

	r.HandleFunc("/", middlewares.Logger(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Backend is up and running!"))
		},
	)).Methods(http.MethodGet)

	api := r.PathPrefix("/api/v1").Subrouter()

	// NOTE Add routes here
	api.HandleFunc("/account/exists", gc.CheckAccountExists).Methods(http.MethodPost)
	api.HandleFunc("/account/create", gc.CreateAccount).Methods(http.MethodPost)

	return r, nil
}
