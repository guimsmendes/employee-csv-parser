package repository

import (
	"employee-csv-parser/configuration"
	"employee-csv-parser/utils"
	"fmt"

	"github.com/hashicorp/go-memdb"
)

type Employee struct {
	Id, FullName, Email, Salary string
}

var db *memdb.MemDB
var txn *memdb.Txn

func (employee *Employee) PersistData() bool {
	db = configuration.GetConnection()

	if !validUnique("id", employee.Id) ||
		!validUnique("email", employee.Email) {
		return false
	}

	txn = db.Txn(true)
	if err := txn.Insert("employee", employee); err != nil {
		fmt.Println(err)
	}
	txn.Commit()
	return true
}

func validUnique(field, value string) bool {
	txn = db.Txn(false)
	defer txn.Abort()
	it, err := txn.Get("employee", field, value)
	utils.PrintError(err)
	if it.Next() != nil {
		fmt.Println(field + " " + value + " already exists.")
		return false
	}
	return true
}
