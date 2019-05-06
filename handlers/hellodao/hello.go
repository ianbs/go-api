package hellodao

import (
	"encoding/json"
	"net/http"
)

//Message type
type Message struct {
	Content string
}

//HelloGet func
func HelloGet(w http.ResponseWriter, r *http.Request) {
	msg := Message{"Hello World from cotroller"}
	json.NewEncoder(w).Encode(msg)
}
