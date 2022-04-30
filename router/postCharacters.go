package router

import (
	"GoAssignment-1/model"
	"GoAssignment-1/operations"
	"encoding/json"
	"net/http"
)
// Function to add characters to the list
func PostCharacters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newCharacter model.FinalList
	_ = json.NewDecoder(r.Body).Decode(&newCharacter)
	List = append(List, newCharacter)
	json.NewEncoder(w).Encode(newCharacter)
	operations.Sorting(List)
	List = operations.RemoveCharacters(List)
}
