package router

import (
	"GoAssignment-1/operations"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Function to get maxpower of specific characters by it's name
func GetCharactersPower(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var flag = 0
	name := vars["name"]
	name = strings.ToLower(name)
	list := GetCharacters()
	for i := 0; i < len(list); i++ {
		if name == list[i].Name {
			flag = 1
			jsonString, _ := json.Marshal(list[i].MaxPower)
			list[i].Count += 1
			fmt.Fprintln(w, string(jsonString))
			break
		}
	}
	if flag == 0 {
		fmt.Fprintf(w, "Not available")
	}
	operations.Sorting(List)
	List = operations.RemoveCharacters(List)
}
