/**
 * Go company application service
 * @author Diego Andre Poli <diegoandrepoli@gmail.com>
 */
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

/**
 * Main execution
 */
func main(){
	serverHandler()
}

/**
 * Create server handler
 */
func serverHandler(){
	var router = mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")

	fmt.Println("Running server server!")
	log.Fatal(http.ListenAndServe(":3030", router))
}

/**
 * Get index message
 * @return service message as string
 */
func getIndexMessage() string {
	return "Welcome to company!"
}

/**
 * Home aplication response message
 */
func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, getIndexMessage())
}