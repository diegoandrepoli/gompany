package repository

import (
	"database/sql"
	"fmt"
	"github.com/diegoandrepoli/gompany/configs"
	"github.com/diegoandrepoli/gompany/model"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

/**
 * Get database connection
 * @return string of database connection
 */
func getConnection() string {

	//get application configuration
	var configs = configs.GetConfiguration()

	//get database configuration
	host := configs.DatabaseConfigHost
    port := configs.DatabaseConfigPort
	user := configs.DatabaseConfigUser
	password := configs.DatabaseConfigPassword
	dbname := configs.DatabaseConfigDbname

	//create database url
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

/**
 * Get database
 * @return string as database type
 */
func getDatabase() string {
	var configs = configs.GetConfiguration()
	return configs.DatabaseConfigDriver
}

/**
 * Save company
 * @param company
 * @retrun company with database id
 */
func SaveCompany(company model.Company) model.Company {
	//open database connection
	db, err := sql.Open(getDatabase(), getConnection())

	//abort on connection error
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//create query
	sqlStatement := `INSERT INTO collector.companies (name, zip, website) VALUES ($1, $2, $3) RETURNING id`

	id := 0

	//execute statement
	err = db.QueryRow(sqlStatement, company.Name, company.Zip, company.Website).Scan(&id)

	//abort on insert error
	if err != nil {

		//is pg error
		if err, ok := err.(*pq.Error); ok {

			//discart as duplicate row
			if err.Code != pq.ErrorCode("23505") {
				panic(err)
			}
		}
	}

	//set company id
	company.ID = id

	return company
}

/**
 * Save company list
 * @param list of company
 */
func SaveCompanyList(companies []model.Company){
	for _, company := range companies {
		SaveCompany(company)
	}
}

/**
 * Update company
 * @param company
 * @return company updated
 */
func UpdateCompany(company model.Company) model.Company {
	//open database connection
	db, err := sql.Open(getDatabase(), getConnection())

	//abort on connection error
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//create query
	sqlStatement := `UPDATE collector.companies set name=$1, zip=$2, website=$3 WHERE id=$4`

	//execute statement
	_, err = db.Exec(sqlStatement, company.Name, company.Zip, company.Website, company.ID)

	//abort on insert error
	if err != nil {
		panic(err)
	}

	return company
}

/**
 * Update company on name based
 * @param company model
 * @return company model updated
 */
func UpdateCompanyByName(company model.Company) model.Company {
	//open database connection
	db, err := sql.Open(getDatabase(), getConnection())

	//abort on connection error
	if err != nil {
		panic(err)
	}

	defer db.Close()

	//create query
	sqlStatement := `UPDATE collector.companies set name=$1, zip=$2, website=$3 WHERE name=$4`

	//execute statement
	_, err = db.Exec(sqlStatement, company.Name, company.Zip, company.Website, company.Name)

	//abort on insert error
	if err != nil {
		panic(err)
	}

	return company
}

/**
 * Update as company by list
 * @param list of company model
 */
func UpdateCompanyByNameList(companies []model.Company){
	for _, company := range companies {
		UpdateCompanyByName(company)
	}
}
