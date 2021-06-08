package repository

import (
	"employee-csv-parser/configuration"
	"testing"
)

func getEmployee() Employee {
	return Employee{Id: "RT1", FullName: "Guilherme", Email: "guimendes@mail.com", Salary: "10"}
}
func TestUniqueData(t *testing.T) {
	configuration.DbConnect()
	employee := getEmployee()
	if employee.PersistData() {
		t.Logf("Success: The data registered was unique. Successfully persisted!")
	} else {
		t.Errorf("Fail: The data registered wasn't unique. Failed to persist!")
	}

}

func TestNotUniqueData(t *testing.T) {
	configuration.DbConnect()
	employee := getEmployee()
	employee.PersistData()
	if !employee.PersistData() {
		t.Logf("Success: The data registered wasn't unique. Failed to persist!")
	} else {
		t.Errorf("Fail: The data registered was unique. Successfully persisted!")
	}

}
