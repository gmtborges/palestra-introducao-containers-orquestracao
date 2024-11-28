package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "mysql"
	}

	connectionString := "root:password@tcp(" + dbHost + ":3306)/testdb"
	return sql.Open("mysql", connectionString)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	db, err := getDB()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		http.Error(w, "Database ping failed", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "healthy", "message": "Database connection successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", healthCheck)
	log.Printf("Listen on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
