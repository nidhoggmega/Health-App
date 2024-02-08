package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	// Connect to MySQL database
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/healthcare_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/patient-records", getPatientRecords).Methods("GET")
	router.HandleFunc("/patient-records", createPatientRecord).Methods("POST")
	router.HandleFunc("/patient-records/{id}", deletePatientRecord).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getPatientRecords(w http.ResponseWriter, r *http.Request) {
	// Perform database query
	rows, err := db.Query("SELECT * FROM patient_records")
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process query results
	var patientRecords []string
	for rows.Next() {
		var patientRecord string
		if err := rows.Scan(&patientRecord); err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		patientRecords = append(patientRecords, patientRecord)
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"patientRecords": patientRecords}); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func createPatientRecord(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var patientRecord string
	err := json.NewDecoder(r.Body).Decode(&patientRecord)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Perform database insert
	_, err = db.Exec("INSERT INTO patient_records (patient_record) VALUES (?)", patientRecord)
	if err != nil {
		log.Println("Error inserting into database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Patient record created successfully")
}

func deletePatientRecord(w http.ResponseWriter, r *http.Request) {
	// Extract patient record ID from request parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Perform database delete
	_, err := db.Exec("DELETE FROM patient_records WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting from database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return response
	fmt.Fprintln(w, "Patient record deleted successfully")
}
