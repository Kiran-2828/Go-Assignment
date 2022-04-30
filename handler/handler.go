package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"GoAssignment-1/router"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", router.GetCharactersList).Methods("GET")
	myRouter.HandleFunc("/{name}", router.GetCharactersPower).Methods("GET")
	myRouter.HandleFunc("/character", router.PostCharacters).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:9090", myRouter))
}