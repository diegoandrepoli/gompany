

package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)


func createCompanyAsLine(){

}

func readCsv(filename string) []Company{
	csvFile, _ := os.Open(filename)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	var companies []Company

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		var website = ""

		if(len(line) > 2){
			website = line[2]
		}

		//apend company
		companies = append(companies, Company{Name: line[0], Zip: line[1], Website: website })

	}

	return companies
}