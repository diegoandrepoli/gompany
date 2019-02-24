
package main

import(

)

// Gompany structure
type Company struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Zip      string   `json:"zip,omitempty"`
	Website  string   `json:"website,omitempty"`
}

func main(){

	//save csv company list
	saveCompanyList(readCsv("q1_catalog.csv"))

	//update csv company list
	updateCompanyByNameList(readCsv("q2_clientData.csv"))
}


