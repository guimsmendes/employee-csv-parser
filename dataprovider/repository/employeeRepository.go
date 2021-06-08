package repository

import (
	"employee-csv-parser/configuration"
	"employee-csv-parser/utils"
	"fmt"
)

type EmployeeEntity struct {
	Id, FullName, Email, Salary string
}

func (employee *EmployeeEntity) PersistData() bool {
	db := configuration.DbConnect()
	fmt.Println(employee)
	selectExistingIdOrEmail, err := db.Query("select id from employee where id = ? or email = ?")
	utils.PrintError(err)
	for selectExistingIdOrEmail.Next() {
		var id string
		err = selectExistingIdOrEmail.Scan(&id)
		utils.PrintError(err)
		if id != "" {
			fmt.Println("Employee with id: " + employee.Id + " and e-mail: " + employee.Email + " already exists in the database.")
			defer db.Close()
			return false
		}
	}

	defer db.Close()
	return true
}
