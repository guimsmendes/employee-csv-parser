package configuration

import (
	"database/sql"
	"employee-csv-parser/utils"
)

func DbConnect() *sql.DB {
	connect := "user=postgres dbname=rain password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connect)
	utils.PrintError(err)
	return db
}
