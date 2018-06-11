package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LocationAPI interface
type LocationAPI interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// Location struct
type Location struct {
	Name string `json:"name"`
}

// Create returns Hello world
func (lg Location) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t Location
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	json.NewEncoder(w).Encode(t)
	fmt.Println(t)
}

// Get retunrs locationGroup
func (lg Location) Get(w http.ResponseWriter, r *http.Request) {
	var res = "test"
	json.NewEncoder(w).Encode(res)
	fmt.Println(res)
}
