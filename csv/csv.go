
package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"github.com/diegoandrepoli/gompany/model"
)

/**
 * Comma, csv separator
 */
var comma = ';'

/**
 * Get csv reader
 * @param name as csv file
 * @return csv file reader
 */
func getCsvReader(filename string) *csv.Reader {
	//open csv file
	csvFile, _ := os.Open(filename)

	//bufferizing
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//set custom coma
	reader.Comma = comma
	return reader
}

/**
 * Read csv file by filemane
 * @param name as csv file
 * @return company list
 */
func ReadCsv(filename string) []model.Company{

	//get file
	reader := getCsvReader(filename)

	//list of csv companies
	var companies []model.Company

	for {
		//read line
		line, error := reader.Read()

		//prepend error
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//append new company
		companies = append(companies, model.CreateCompanyByArray(line))
	}

	return companies
}