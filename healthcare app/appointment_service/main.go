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
	router.HandleFunc("/appointments", getAppointments).Methods("GET")
	router.HandleFunc("/appointments", createAppointment).Methods("POST")
	router.HandleFunc("/appointments/{id}", deleteAppointment).Methods("DELETE")

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAppointments(w http.ResponseWriter, r *http.Request) {
	// Perform database query
	rows, err := db.Query("SELECT * FROM appointments")
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process query results
	var appointments []string
	for rows.Next() {
		var appointment string
		if err := rows.Scan(&appointment); err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		appointments = append(appointments, appointment)
	}

	// Return response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"appointments": appointments}); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func createAppointment(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var appointment string
	err := json.NewDecoder(r.Body).Decode(&appointment)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Perform database insert
	_, err = db.Exec("INSERT INTO appointments (appointment) VALUES (?)", appointment)
	if err != nil {
		log.Println("Error inserting into database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Appointment created successfully")
}

func deleteAppointment(w http.ResponseWriter, r *http.Request) {
	// Extract appointment ID from request parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Perform database delete
	_, err := db.Exec("DELETE FROM appointments WHERE id = ?", id)
	if err != nil {
		log.Println("Error deleting from database:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return response
	fmt.Fprintln(w, "Appointment deleted successfully")
}
