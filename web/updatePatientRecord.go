package web

import (
	//"database/sql"
	"fmt"
	"Sano/database"

	_ "github.com/mattn/go-sqlite3"
)

func updatePatientRecord(patientID string, driveFileID string) error {

	stmt, err := db.DB.Prepare("UPDATE patients SET drive_file_id = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("unable to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(driveFileID, patientID)
	if err != nil {
		return fmt.Errorf("unable to execute SQL statement: %v", err)
	}

	return nil
}
