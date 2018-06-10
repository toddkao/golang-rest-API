package locations

import (
	"encoding/json"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello World")
}
