package usecase

import (
	"regexp"
)

func validContent(value string, pattern string) bool {

	validator := regexp.MustCompile(pattern)
	if validator.MatchString(value) {
		return true
	}
	return false
}
