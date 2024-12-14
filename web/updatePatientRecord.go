package web

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Assume you have a global variable for the database connection
var db *sql.DB

// Initialize the database connection (this should be done once, e.g., in main or init function)
func init() {
	var err error
	// Open a connection to the SQLite database file
	db, err = sql.Open("sqlite3", "./mydb.sqlite")
	if err != nil {
		panic(err)
	}
}

// updatePatientRecord updates the patient's record with the new file ID
func updatePatientRecord(patientID string, driveFileID string) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("UPDATE patients SET drive_file_id = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("unable to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(driveFileID, patientID)
	if err != nil {
		return fmt.Errorf("unable to execute SQL statement: %v", err)
	}

	return nil
}
