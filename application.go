package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Gompany structure
type Company struct {
	ID       int   `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Zip      string   `json:"zip,omitempty"`
	Website  string   `json:"website,omitempty"`
}

/**
 * Main server
 */
func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/company", postCompanyApi).Methods("POST")

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/**
 * Main gompany action
 */
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, welcome to Gompany!")
}


/**
 * Post company api
 */
func postCompanyApi(w http.ResponseWriter, r *http.Request) {
	var company Company

	//decode api content
	_ = json.NewDecoder(r.Body).Decode(&company)

	//save and return company
	json.NewEncoder(w).Encode(saveCompany(company))
}


