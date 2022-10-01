package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetVars(w http.ResponseWriter, r *http.Request) map[string]string {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	return vars
}
