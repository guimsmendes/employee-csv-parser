package usecase

import (
	"fmt"
	"strings"
)

func VerifyTable(employeeTable [][]string) bool {

	emailPosition := getEmailPosition(employeeTable)
	fmt.Println(emailPosition)
	return true

}

func getEmailPosition(employeeTable [][]string) int {

	line := employeeTable[0]

	for i := 0; i < len(line); i++ {
		if strings.Contains(strings.ToLower(line[i]), "mail") {
			validEmailContent(employeeTable, i)
			return i
		}
	}
	return -1
}

func validEmailContent(employeeTable [][]string, position int) {

	for i := 1; i < len(employeeTable); i++ {
		fmt.Println(employeeTable[i][position])
	}

}
