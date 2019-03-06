
package main

import (
	"github.com/diegoandrepoli/gompany/repository"
	"github.com/diegoandrepoli/gompany/csv"
	"github.com/diegoandrepoli/gompany/configs"
)

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


