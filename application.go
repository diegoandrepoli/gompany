package main

import (
	"encoding/json"
	"fmt"
	"github.com/diegoandrepoli/gompany/model"
	"github.com/diegoandrepoli/gompany/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

/**
 * Main server
 */
func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/company", saveCompanyApi).Methods("POST")
	router.HandleFunc("/company", updateCompanyApi).Methods("PUT")

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
func saveCompanyApi(w http.ResponseWriter, r *http.Request) {
	var company model.Company

	//decode api content
	_ = json.NewDecoder(r.Body).Decode(&company)


	//save and return company
	json.NewEncoder(w).Encode(repository.SaveCompany(company))
}

/**
 * Put as update company
 */
func updateCompanyApi(w http.ResponseWriter, r *http.Request) {
	var company model.Company

	//decode api content
	_ = json.NewDecoder(r.Body).Decode(&company)

	//save and return company
	json.NewEncoder(w).Encode(repository.UpdateCompany(company))
}


