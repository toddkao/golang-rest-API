package location

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LocationGroupAPI defines all operations
type LocationGroupAPI interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

// LocationGroup returns locationGroup Struct
type LocationGroup struct {
	Name string `json:"name"`
}

// Create returns Hello world
func (lg *LocationGroup) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t LocationGroup
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	json.NewEncoder(w).Encode(t)
	fmt.Println(t)
}

// Get retunrs locationGroup
func (lg *LocationGroup) Get(w http.ResponseWriter, r *http.Request) {
	var res = "test"
	json.NewEncoder(w).Encode(res)
	fmt.Println(res)
}
