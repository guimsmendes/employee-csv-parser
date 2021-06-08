package domain

import (
	"fmt"
	"strings"
)

const salaryPattern string = "[$]{0,1}[ ]{0,2}[0-9]{1,}([.][0-9]{1,2}\b){0,1}"

func GetSalaryPosition(line []string) int {

	for i, value := range line {
		if strings.Contains(strings.ToLower(value), "wage") ||
			strings.Contains(strings.ToLower(value), "sal") ||
			strings.Contains(strings.ToLower(value), "rate") {
			fmt.Println("Salary column successfuly retrieved!")
			return i
		}

	}
	err := fmt.Errorf("Salary column not found on this file.")
	fmt.Println(err.Error())
	return -1
}

func GetSalaryPattern() string {
	return salaryPattern
}
