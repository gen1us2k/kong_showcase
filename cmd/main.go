package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloJSON(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Greetings!"}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func main() {
	http.HandleFunc("/greet", helloJSON)
	http.ListenAndServe(":8090", nil)

}
