package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

/**
 * Database configuration
 */
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

/**
 * Get database connection
 * @return string of database connection
 */
func getConnection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

/**
 * Get database
 * @return string as database type
 */
func getDatabase() string {
	return "postgres"
}

/**
 * Save company
 * @param company
 * @retrun company with database id
 */
func saveCompany(company Company) Company {
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
		panic(err)
	}

	//set company id
	company.ID = id

	return company
}
