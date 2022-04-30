package router

import (
	"net/http"
	"encoding/json"
)
//Function for displaying the characters and their information
func GetCharactersList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(List)
}

