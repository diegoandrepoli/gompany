
package main

import (
	"github.com/diegoandrepoli/gompany/repository"
	"github.com/diegoandrepoli/gompany/csv"
	"github.com/diegoandrepoli/gompany/configs"
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
	var configs = configs.GetConfiguration()

	//save csv company list
	repository.SaveCompanyList(csv.ReadCsv(configs.FileCatalogName))

	//update csv company list
	repository.UpdateCompanyByNameList(csv.ReadCsv(configs.FileClientName))
}


