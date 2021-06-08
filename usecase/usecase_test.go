package usecase

import (
	"employee-csv-parser/configuration"
	"employee-csv-parser/usecase/domain"
	"testing"
)

func getValidTable() [][]string {
	return [][]string{{"E-mail", "Name", "Id", "Wage"}, {"guimendes@mail.com", "Guilherme Mendes", "RT4", "10.5"}}
}

func getInvalidTable() [][]string {
	return [][]string{{"Size", "Shape"}}
}

func TestValidTable(t *testing.T) {

	configuration.DbConnect()

	validator := VerifyTable(getValidTable())
	if validator {
		t.Logf("Success: The Validation expected for verifying the invalid Table was true, got %v", validator)
	} else {
		t.Errorf("Fail: The Validation expected for verifying the invalid Table was true, got %v", validator)
	}

}

func TestInvalidTable(t *testing.T) {

	validator := VerifyTable(getInvalidTable())
	if validator {
		t.Errorf("Fail: The Validation expected for verifying the invalid Table was false, got %v", validator)
	} else {
		t.Logf("Success: The Validation expected for verifying the invalid Table was false, got %v", validator)
	}
}

func TestValidContent(t *testing.T) {

	validator := validContent(getValidTable()[1][0], domain.GetEmailPattern())

	if validator {
		t.Logf("Success: The Validation expected for verifying the content pattern was true, got %v", validator)
	} else {
		t.Errorf("Fail: The Validation expected for verifying the content pattern was true, got %v", validator)
	}

}

func TestInvalidContent(t *testing.T) {

	validator := validContent(getValidTable()[1][0], domain.GetNamePattern())

	if !validator {
		t.Logf("Success: The Validation expected for verifying the content pattern was false, got %v", validator)
	} else {
		t.Errorf("Fail: The Validation expected for verifying the content pattern was false, got %v", validator)
	}

}
