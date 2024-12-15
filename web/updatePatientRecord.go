package web

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/mattn/go-sqlite3"
// )

// var db *sql.DB

// func init() {
// 	var err error
	
// 	db, err = sql.Open("sqlite3", "./mydb.sqlite")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func updatePatientRecord(patientID string, driveFileID string) error {

// 	stmt, err := db.Prepare("UPDATE patients SET drive_file_id = ? WHERE id = ?")
// 	if err != nil {
// 		return fmt.Errorf("unable to prepare SQL statement: %v", err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(driveFileID, patientID)
// 	if err != nil {
// 		return fmt.Errorf("unable to execute SQL statement: %v", err)
// 	}

// 	return nil
// }
