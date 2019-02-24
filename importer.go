
package main

import (

)

// Gompany structure
type Company struct {
	ID       int      `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Zip      string   `json:"zip,omitempty"`
	Website  string   `json:"website,omitempty"`
}

/**
 * Main importer
 * Execute importer from file catalog name and client file name
 */
func main(){

	//get application configurations
	var configs = getConfiguration()

	//save csv company list
	saveCompanyList(readCsv(configs.FileCatalogName))

	//update csv company list
	updateCompanyByNameList(readCsv(configs.FileClientName))
}


