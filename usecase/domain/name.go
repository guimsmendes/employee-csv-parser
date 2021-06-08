package domain

import (
	"fmt"
	"strings"
)

const namePattern string = "[A-Z][a-z]{1,15}"
const fullNamePattern string = "[A-Z][a-z]{1,15}[ ][A-Z][a-z]{1,15}"

func GetNamePosition(line []string) (name int, first int, last int) {
	var count int
	for i, value := range line {

		if strings.Contains(strings.ToLower(value), "name") {
			count++
			name = i
		}
		if strings.Contains(strings.ToLower(value), "first") || strings.Contains(strings.ToLower(value), "f.") {
			first = i
		}
		if strings.Contains(strings.ToLower(value), "last") || strings.Contains(strings.ToLower(value), "l.") {
			last = i
		}

	}

	if name == 0 && first == 0 && last == 0 {
		err := fmt.Errorf("Name column not found on this file.")
		fmt.Println(err.Error())
	}

	switch count {
	case 1:
		fmt.Println("Full name column successfully retrieved!")
		return name, 0, 0
	default:
		fmt.Println("First and Last name columns successfully retrieved!")
		return 0, first, last
	}
}

func GetNamePattern() string {
	return namePattern
}

func GetFullNamePattern() string {
	return fullNamePattern
}
