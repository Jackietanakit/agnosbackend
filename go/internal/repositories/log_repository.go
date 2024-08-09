package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

func StoreLog(request map[string]string, response map[string]interface{}) {
	db, err := sql.Open("postgres", "user=admin password=qwerqwer dbname=agnos_db sslmode=disable")
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
