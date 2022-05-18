package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"miproyecto/numbers"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println(numbers.IsEven(3))
	r := mux.NewRouter()
	r.HandleFunc("/api/greetings/{name}", func(w http.ResponseWriter, r *http.Request) {
		m := mux.Vars(r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello " + m["name"])
	}).Methods(http.MethodGet)
	r.HandleFunc("/api/helloWorld", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello World")
	}).Methods(http.MethodGet)

	fmt.Println("The server is up on port 3000")
	http.ListenAndServe(":3000", r)

}
