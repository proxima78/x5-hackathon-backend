package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/proxima78/x5-hackathon-backend/app"
	"github.com/proxima78/x5-hackathon-backend/middlewares"
)

func NewRouter(a *app.App) (*mux.Router, error) {
	r := mux.NewRouter()

	// NOTE Create repositories here

	// NOTE Create services here

	// NOTE Create controllers here

	// NOTE Create middlewares here

	r.HandleFunc("/", middlewares.Logger(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Backend is up and running!"))
		},
	)).Methods(http.MethodGet)

	//api := r.PathPrefix("/api/v1").Subrouter()

	// NOTE Add routes here

	return r, nil
}
