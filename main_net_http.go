package main

import (
	"encoding/json"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	var person = struct {
		Name string
		Age  int
	}{
		Name: "golang-gin", Age: 8,
	}
	//  application/json
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&person)
	// io.WriteString(w, "xxxx")
	// fmt.Fprintf(w , "xxxx")
}

func main_net_http() {

	http.HandleFunc("/welcome", welcome)

	http.ListenAndServe(":8080", nil)
}
