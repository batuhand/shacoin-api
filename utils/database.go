package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {

	connStr := fmt.Sprintf("host= %s port= %d user=%s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	Db, _ = sql.Open("postgres", connStr)

}
