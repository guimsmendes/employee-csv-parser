package domain

import (
	"fmt"
	"strings"
)

const idPattern string = "[A-Z]{2}[0-9]{1,3}"

func GetIdPosition(line []string) int {

	for i, value := range line {
		if strings.Contains(strings.ToLower(value), "id") ||
			(strings.Contains(strings.ToLower(value), "number") &&
				strings.Contains(strings.ToLower(value), "emp")) {
			fmt.Println("Id column successfuly retrieved!")
			return i
		}

	}
	err := fmt.Errorf("Id column not found on this file.")
	fmt.Println(err.Error())
	return -1
}

func GetIdPattern() string {
	return idPattern
}
