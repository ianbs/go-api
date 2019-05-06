package persons

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var persons = []Person{
	Person{ID: "1", Name: "Ian", Age: 27},
	Person{ID: "2", Name: "John", Age: 45},
}

//Message type for return a message
type Message struct {
	Content string
	Found   bool
}

//GetAll persons function get all persons
func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(persons)
}

//GetOne persons function show one person
func GetOne(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range persons {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Message{Content: "Not found", Found: false})
}
