
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
 * Comma - csv file separator
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
		companies = append(companies, createCompanyByLine(line))
	}

	return companies
}

/**
 * Get line reader by index
 * @param list of items - csv line
 * @param index of item
 * @return value as item or empty is out of range
 */
func getLineReaderByIndex(items []string, index int) string{
	if(len(items) > index){
		return items[index]
	}

	return ""
}

/**
 * Create company by line
 * @param list of items - csv line
 * @return company
 */
func createCompanyByLine(items []string) model.Company {
	var company model.Company

	//populate company
	company.Name = getLineReaderByIndex(items, 0)
	company.Zip = getLineReaderByIndex(items, 1)
	company.Website = getLineReaderByIndex(items, 2)

	return company
}




