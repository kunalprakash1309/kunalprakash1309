// When we mention the package names, all the
// Go programs in that package can share variables and functions exported without any need
// of actual importing.


package dbutils

import (
	"log"
	"database/sql"
)

func Initialize(dbDriver *sql.DB) {
	// for train table
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already Exists or error in creating train table")
	}

	// for station table
	statement, driverError = dbDriver.Prepare(station)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create station table
	_, statementError = statement.Exec()
	if statementError != nil {
		log.Println("Table already Exists or error in creating station table")
	}

	// for schedule table
	statement, driverError = dbDriver.Prepare(schedule)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create schedule table
	_, statementError = statement.Exec()
	if statementError != nil {
		log.Println("Table already Exists or error in creating schedule table")
	}

	log.Println("All tables created/initialized successfully")
}