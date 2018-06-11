package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/toddkao/ecommGo/internal/app/routes/location"
)

// InitRoutes to initialize API routes
func InitRoutes() {
	router := mux.NewRouter()
	var locationGroup LocationGroup
	var location Location

	router.HandleFunc("/locationGroups/create", locationGroup.Create).Methods("POST")
	router.HandleFunc("/locationGroups", locationGroup.Get).Methods("GET")
	router.HandleFunc("/location/create", location.Create).Methods("POST")
	router.HandleFunc("/location", location.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":7777", router))
}
