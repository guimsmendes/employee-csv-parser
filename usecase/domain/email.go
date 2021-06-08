package domain

import (
	"fmt"
	"strings"
)

const emailPattern string = ".{2,}@[a-z]{2,}([.][a-z]{2,3}){1,3}"

func GetEmailPosition(line []string) int {

	for i, value := range line {
		if strings.Contains(strings.ToLower(value), "mail") {
			fmt.Println("E-mail column successfuly retrieved!")
			return i
		}
	}
	err := fmt.Errorf("Name column not found on this file.")
	fmt.Println(err.Error())
	return -1
}

func GetEmailPattern() string {
	return emailPattern
}
