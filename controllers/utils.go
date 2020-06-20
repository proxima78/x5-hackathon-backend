package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnHTTPError(w http.ResponseWriter, err string, code int) {
	res := httpModels.HTTPError{
		Error: models.HTTPErrorMessage{
			Message: err,
		},
	}
	r, _ := json.Marshal(res)
	writeHTTPJSONError(w, string(r), code)
}

func writeHTTPJSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, err)
}
