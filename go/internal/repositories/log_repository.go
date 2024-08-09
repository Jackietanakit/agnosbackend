package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func StoreLog(request map[string]string, response map[string]interface{}) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	reqJSON, _ := json.Marshal(request)
	resJSON, _ := json.Marshal(response)

	_, err = db.Exec("INSERT INTO logs (request_body, response_body) VALUES ($1, $2)", string(reqJSON), string(resJSON))
	if err != nil {
		fmt.Println(err)
	}
}
