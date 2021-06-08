package usecase

import (
	"employee-csv-parser/dataprovider"
	"employee-csv-parser/dataprovider/repository"
	"employee-csv-parser/usecase/domain"
	"fmt"
	"strconv"
	"strings"
)

var validLine int
var invalidLine int

var emailPosition int
var salaryPosition int
var idPosition int
var fullName int
var firstName int
var lastName int

func VerifyTable(employeeTable [][]string) bool {

	validData, invalidData := initializeData(employeeTable)
	setPosition(employeeTable[0])

	if emailPosition == -1 || salaryPosition == -1 || idPosition == -1 {
		for i := 1; i < len(employeeTable); i++ {
			writeInvalidLine(employeeTable[i], invalidData)
		}
		return false
	}

	for i := 1; i < len(employeeTable); i++ {
		if validContent(employeeTable[i][emailPosition], domain.GetEmailPattern()) &&
			validContent(employeeTable[i][idPosition], domain.GetIdPattern()) &&
			validContent(employeeTable[i][salaryPosition], domain.GetSalaryPattern()) &&
			(validContent(employeeTable[i][fullName], domain.GetFullNamePattern()) ||
				(validContent(employeeTable[i][firstName], domain.GetNamePattern()) &&
					validContent(employeeTable[i][lastName], domain.GetNamePattern()))) {

			fmt.Println("Line " + strconv.Itoa(i) + " is valid: " + strings.Join(employeeTable[i], " "))

			writeValidLine(employeeTable[i], validData)
		} else {
			err := fmt.Errorf("Line " + strconv.Itoa(i) + " is invalid:" + strings.Join(employeeTable[i], " "))
			fmt.Println(err)
			writeInvalidLine(employeeTable[i], invalidData)
		}
	}

	persist(validData)
	fmt.Println("Registering Invalid Data on csv file.")
	for i, line := range invalidData {
		if strings.Join(line, "") == "" {
			break
		}
		fmt.Println("Registering Invalid Data for line " + strconv.Itoa(i+1) + " on csv file.")
		dataprovider.RegisterCsv(line, "invalid")
	}
	return true
}

func initializeData(employeeTable [][]string) ([][]string, [][]string) {

	validLine = 0
	invalidLine = 0
	validData := make([][]string, len(employeeTable))
	invalidData := make([][]string, len(employeeTable))
	for i := 0; i < len(employeeTable); i++ {
		validData[i] = make([]string, 4)
		invalidData[i] = make([]string, len(employeeTable[i]))
	}

	return validData, invalidData
}

func setPosition(line []string) {

	emailPosition = domain.GetEmailPosition(line)
	salaryPosition = domain.GetSalaryPosition(line)
	idPosition = domain.GetIdPosition(line)
	fullName, firstName, lastName = domain.GetNamePosition(line)

}

func writeValidLine(line []string, validData [][]string) {
	var name string
	if fullName != 0 {
		name = line[fullName]
	} else {
		name = line[firstName] + " " + line[lastName]
	}
	validData[validLine] = []string{line[idPosition], name, line[emailPosition], line[salaryPosition]}
	validLine++

}

func writeInvalidLine(line []string, invalidData [][]string) {
	invalidData[invalidLine] = line
	invalidLine++
}

func persist(validData [][]string) {
	for i, line := range validData {
		if strings.Join(line, "") == "" {
			break
		}
		employee := repository.Employee{Id: line[0], FullName: line[1], Email: line[2], Salary: line[3]}
		fmt.Println("Line " + strconv.Itoa(i+1) + " is being persisted. Employee Id: " + employee.Id)
		if employee.PersistData() {
			fmt.Println("Line " + strconv.Itoa(i+1) + " for Employee Id: " + employee.Id + " successfully persisted.")
			fmt.Println("Registering Valid Data for line " + strconv.Itoa(i+1) + " on csv file.")
			dataprovider.RegisterCsv(line, "valid")
		}
	}
}
